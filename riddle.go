package main

type Riddle struct {
	C int      `json:"chars"`
	P []string `json:"participants"`
	W string   `json:"word"` // TODO
}

func NewRiddle(r *Room) *Riddle {
	riddle := Riddle{}
	riddle.C = len(r.W)
	riddle.P = make([]string, len(r.participants))
	for i := 0; i < len(r.participants); i++ {
		p := r.participants[i]
		riddle.P = append(riddle.P, p.name)
	}
	riddle.W = room.W
	return &riddle
}
