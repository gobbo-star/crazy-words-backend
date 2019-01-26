package main

type Riddle struct {
	C int            `json:"chars"`
	P []*Participant `json:"participants"`
	W string         `json:"word"` // TODO
}

func NewRiddle(r *Room) *Riddle {
	riddle := Riddle{}
	riddle.C = len(r.W)
	riddle.P = r.participants
	riddle.W = room.W
	return &riddle
}
