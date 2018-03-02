package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rng = struct {
	rand *rand.Rand
	// Mutex is needed just for safety since rand.NewSource
	// is not safe for concurrent use.
	sync.Mutex
}{}

func init() {
	rng.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {

	for i := 0; i < 10; i++ {
		rng.Lock()
		got := rng.rand.Intn(6)
		rng.Unlock()
		fmt.Println(got)
	}
}
