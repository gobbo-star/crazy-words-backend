package main

import (
	"math/rand"
	"time"
)

type NameGen struct {
	rnd *rand.Rand
	cg  *ColorGen
}

func NewNameGen(cg *ColorGen) *NameGen {
	ng := NameGen{}
	ng.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	ng.cg = cg
	return &ng
}

func (ng *NameGen) GenName() string {
	return ng.cg.GenColor().name
}
