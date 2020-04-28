package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	var Ball int

	fmt.Printf("Start Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("Start Number of Go routines: %d\n", runtime.NumGoroutine())

	tableCh := make(chan int)
	fmt.Printf("[MAIN GID:%d]\n", getGID())
	for i := 0; i < 20; i++ {
		go player(tableCh)
	}

	fmt.Printf("During the process | Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("During the process | Number of Go routines: %d\n", runtime.NumGoroutine())

	tableCh <- Ball
	time.Sleep(10 * time.Second)
	close(tableCh)

	fmt.Printf("End the process | Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("End the process | Number of Go routines: %d\n", runtime.NumGoroutine())
}

func player(tableCh chan int) {
	for {
		ball := <-tableCh
		ball++
		fmt.Printf("[GID:%d, value:%d]\n", getGID(), ball)
		time.Sleep(100 * time.Millisecond)
		tableCh <- ball
	}
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
