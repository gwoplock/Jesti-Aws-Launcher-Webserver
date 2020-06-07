package webserv

import (
	"io"
	"net/http"
	"strings"

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
	for {
		_, p, err := c.NextReader()
		if err != nil {
			c.Close()
			break
		}

		buf := new(strings.Builder)
		io.Copy(buf, p) // TODO handle error
		logrus.Trace("go WS message: ", buf.String())
	}
}
