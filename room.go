package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Room struct {
	ticker       *time.Ticker
	W            string
	participants map[string]*Participant
	wg           *WordGen
	scoreLimit   uint
}

func (r *Room) Start() {
	for t := range r.ticker.C {
		r.W = r.wg.GenWord()
		fmt.Println(t)
		fmt.Println(r.W)
	}
}

func (r *Room) Join(p *Participant) {
	r.participants[p.Name] = p
}

func (r *Room) Notify() {
	rs, err := json.Marshal(NewRiddle(r))
	for _, p := range r.participants {
		if err != nil {
			fmt.Println(err)
		}
		p.Notify(rs)
	}
}

func (r *Room) Quit(p *Participant) {
	delete(r.participants, p.Name)
}

func (r *Room) Guess(bytes []byte) bool {
	success := r.W == string(bytes)
	return success
}

func (r *Room) NewWord() {
	r.W = r.wg.GenWord()
}

func (r *Room) MaxScoreReached(score uint) bool {
	return score >= r.scoreLimit
}

func (r *Room) NewGame() {
	for _, p := range r.participants {
		p.Score = 0
	}
}

func NewRoom(refreshRate time.Duration, gen *WordGen, scoreLimit uint) *Room {
	r := Room{}
	rand.Seed(time.Now().UnixNano())
	genWordsPool()
	r.wg = gen
	r.W = r.wg.GenWord()
	r.participants = map[string]*Participant{}
	r.ticker = time.NewTicker(refreshRate)
	r.scoreLimit = scoreLimit
	return &r
}
