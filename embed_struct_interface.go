package main

import (
	"log"
)

type blah interface {
	Do()
}

type nah struct {
}

func (n nah) Do() {
	log.Println("yes")
}

type hah struct {
	nah
}

func Do(b blah) {
	b.Do()
}

func main() {
	x := hah{
		nah: nah{},
	}
	Do(x)
}
