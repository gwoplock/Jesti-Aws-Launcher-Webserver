package webserv

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true //TODO remove
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Error("unable to upgrade websocket: ", err)
	}

	wg.Add(1)
	go readLoop(conn)
}

func readLoop(c *websocket.Conn) {
	defer wg.Done()
	defer c.Close()
	for {

		mType, bCont, err := c.ReadMessage()
		if err != nil {
			logrus.Error("Error reading message: ", err)
			return
		}

		if mType == websocket.TextMessage {
			logrus.Trace("go WS message: ", string(bCont))
			var message wsMessage
			err := json.Unmarshal(bCont, &message)
			if err != nil {
				logrus.Error("Error unmarshaling Json: ", err)
			}
			handleMessage(message, c)

		} else {
			logrus.Error("Unexpected message type: ", mType)
		}

	}
}
