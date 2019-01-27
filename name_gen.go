package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type NameGen struct {
	colors []Color
}

type Color struct {
	code string
	name string
}

func NewNameGen(file string) *NameGen {
	ng := NameGen{}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(content), "\n")
	ng.colors = make([]Color, len(lines))
	for i := 0; i < len(lines); i++ {
		ls := strings.Split(lines[i], "\t")
		ng.colors = append(ng.colors, Color{ls[0], ls[1]})
	}
	return &ng
}

func (ng *NameGen) GenName() string {
	return newWordRandLen()
}
