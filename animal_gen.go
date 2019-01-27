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

func (ag AnimalGen) GenAnimal() string {
	return ag.animals[ag.rnd.Intn(len(ag.animals))]
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
		ls := strings.Split(lines[i], "\t")
		ag.animals[i] = ls[0]
	}
	return &ag
}
