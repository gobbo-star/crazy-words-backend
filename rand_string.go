package main

import (
	"math/rand"
	"strings"
)

func randLen() int {
	return rand.Intn(5) + 3
}

func newWordRandLen() string {
	return newWord(randLen())
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
