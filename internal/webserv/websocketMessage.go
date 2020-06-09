package webserv

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func handleMessage(message wsMessage, c *websocket.Conn) {
	switch message.Type {
	case LOG_IN:
		logrus.Trace("got pin: ", message.Content)
		if message.Content == 1153 {
			logrus.Info("pin accepted, starting EC2 instance")
			//TODO start EC2
			c.WriteMessage(websocket.TextMessage, []byte("accepted")) //todo real json message
		} else {
			logrus.Error("Incorrect pin")
			c.WriteMessage(websocket.TextMessage, []byte("rejected")) //todo real json message
		}
		break
	case USER_COUNT:
		break
	}
}
