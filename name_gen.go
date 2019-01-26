package main

type NameGen struct {
}

func NewNameGen() *NameGen {
	ng := NameGen{}
	return &ng
}

func (ng *NameGen) GenName() string {
	return newWordRandLen()
}
