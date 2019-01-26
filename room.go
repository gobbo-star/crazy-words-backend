package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Room struct {
	ticker *time.Ticker
	w      string
}

func (r Room) start() {
	for t := range r.ticker.C {
		r.w = newWord(randLen())
		fmt.Println(t)
	}
}

func NewRoom(refreshRate time.Duration) Room {
	r := Room{}
	rand.Seed(time.Now().UnixNano())
	genWordsPool()
	r.w = newWord(randLen())
	r.ticker = time.NewTicker(refreshRate)
	return r
}

func randLen() int {
	return rand.Intn(5) + 3
}

func newWord(wLen int) string {
	w := strings.Builder{}
	for ; wLen > 0; wLen-- {
		w.WriteRune(runeSet[rand.Intn(len(runeSet))])
	}
	return w.String()
}

func genWordsPool() {
	for r := 'a'; r <= 'z'; r++ {
		runeSet = append(runeSet, r)
	}
}