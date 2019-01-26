package main

import "github.com/gorilla/websocket"

type Participant struct {
	ws   *websocket.Conn
	name string
}

func (p *Participant) Notify(bytes []byte) {
	_ = p.ws.WriteMessage(websocket.TextMessage,
		bytes)
}

func NewParticipant(ws *websocket.Conn, name string) *Participant {
	p := Participant{ws, name}
	return &p
}
