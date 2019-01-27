package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type AnimalGen struct {
	rnd     *rand.Rand
	animals []string
}

func NewAnimalGen(file string) *AnimalGen {
	ag := AnimalGen{}
	ag.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(content), "\n")
	ag.animals = make([]string, len(lines))
	for i := 0; i < len(lines); i++ {
		ag.animals[i] = lines[0]
	}
	return &ag
}
