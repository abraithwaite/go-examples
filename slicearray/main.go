package main

import (
	"bufio"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"reflect"
	"runtime"
	"sync"
	"unsafe"
)

type queue struct {
	sync.Mutex
	fifo []int
}

func (q *queue) push(s int) {
	q.Lock()
	q.fifo = append(q.fifo, s)
	q.Unlock()
}

func (q *queue) pop() (s int) {
	q.Lock()
	if len(q.fifo) == 0 {
		return 0
	}
	q.fifo, s = q.fifo[1:], q.fifo[0]
	q.Unlock()
	return s
}

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	x := queue{}
	pSlice(&x.fifo)
	// Allocate roughly 1 GB
	x.fifo = make([]int, 0, 125000000)
	for i := 0; i < 125000000; i++ {
		x.push(1)
	}
	pSlice(&x.fifo)
	x.fifo = x.fifo[120000000:]
	pSlice(&x.fifo)
	PrintMemUsage()
	fmt.Print("press any key to continue")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	y := make([]int, len(x.fifo))
	copy(y, x.fifo)
	x.fifo = nil
	pSlice(&x.fifo)
	pSlice(&y)
	runtime.GC()
	PrintMemUsage()
}

func pSlice(s *[]int) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(s))
	dataptr := (*int)(unsafe.Pointer(hdr.Data))
	fmt.Printf("len: %d, cap: %d, addr: %p\n", len(*s), cap(*s), dataptr)
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
