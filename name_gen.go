package main

import (
	"fmt"
	"math/rand"
	"time"
)

type NameGen struct {
	rnd *rand.Rand
	cg  *ColorGen
	ag  *AnimalGen
}

func NewNameGen(cg *ColorGen, ag *AnimalGen) *NameGen {
	ng := NameGen{}
	ng.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	ng.cg = cg
	ng.ag = ag
	return &ng
}

func (ng *NameGen) GenName() string {
	return fmt.Sprintf("%v %v", ng.cg.GenColor().name, ng.ag.GenAnimal())
}
