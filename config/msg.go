package config

type Message struct {
	CheckRequest string
}

var Msg *Message

func init() {
	Msg = &Message{
		CheckRequest: "Bad Request",
	}
}
