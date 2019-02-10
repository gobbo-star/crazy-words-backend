package main

type MsgType int

const (
	STATE MsgType = iota
	GUESS
	CHAT
)

type Message struct {
	Payload interface{} `json:"Payload"`
	Type    MsgType
}
