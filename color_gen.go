package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type ColorGen struct {
	colors []Color
	rnd    *rand.Rand
}

type Color struct {
	code string
	name string
}

func NewColorGen(file string) *ColorGen {
	cg := ColorGen{}
	cg.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(content), "\n")
	cg.colors = make([]Color, len(lines))
	for i := 0; i < len(lines); i++ {
		ls := strings.Split(lines[i], "\t")
		if len(lines) < 2 {
			fmt.Printf("line %v (%v) has NOT enough columns in it to present a color\n", i, lines[i])
		}
		cg.colors[i] = Color{ls[1], ls[0]}
	}
	return &cg
}

func (ng *ColorGen) GenColor() Color {
	return ng.colors[ng.rnd.Intn(len(ng.colors))]
}
