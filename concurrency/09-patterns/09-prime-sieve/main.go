package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Printf("[GID:%d,PRIME:%d, ITER: %d, GOROUTINES:%d]\n", getGID(), prime, i, runtime.NumGoroutine())
		//fmt.Printf("During the process | Number of CPUS: %d\n", runtime.NumCPU())
		//fmt.Printf("During the process | Number of Go routines: %d\n", runtime.NumGoroutine())
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		//fmt.Printf("[ch:%#v, ch1:%d]\n", ch, <-ch1)
		ch = ch1
		//fmt.Printf("[ch:%d, ch1:%d]\n", <-ch, <-ch1)
	}
}

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i //Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		fmt.Printf("[in: %d, operation_result:%d]\n", i, i%prime)
		if i%prime != 0 {
			//out <- i
		}
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
