package main

import "github.com/gorilla/websocket"

type Participant struct {
	ws    *websocket.Conn
	Name  string `json:"Name"`
	Score uint   `json:"Score"`
}

func (p *Participant) Notify(bytes []byte) {
	_ = p.ws.WriteMessage(websocket.TextMessage,
		bytes)
}

func NewParticipant(ws *websocket.Conn, name string) *Participant {
	p := Participant{ws, name, 0}
	return &p
}
