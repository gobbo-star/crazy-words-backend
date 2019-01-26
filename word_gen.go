package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type WordGen struct {
	words []string
	rnd   *rand.Rand
}

func (gen WordGen) GenWord() string {
	ix := gen.rnd.Intn(len(gen.words))
	return gen.words[ix]
}

func NewWordGen(file string) *WordGen {
	wg := WordGen{}
	wg.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	wg.words = strings.Split(string(content), "\n")
	return &wg
}
