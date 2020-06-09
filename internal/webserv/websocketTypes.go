package webserv

type wsMessage struct {
	Type    wsMessageType `json:"type"`
	Content int           `json:"content"`
}

type wsMessageType int

const (
	LOG_IN wsMessageType = iota
	USER_COUNT
)
