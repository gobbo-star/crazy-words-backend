package main

type MsgType int

const (
	STATE MsgType = iota
	GUESS
	CHAT
)

var MsgTypes = map[MsgType]string{
	STATE: "STATE",
	GUESS: "GUESS",
	CHAT:  "CHAT",
}

type Message struct {
	Payload interface{} `json:"Payload"`
	Type    string      `json:"Type"`
}
