package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	in := inputGen()
	fmt.Printf("Start Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("Start Number of Go routines: %d\n", runtime.NumGoroutine())

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute  work across multiple functions (then goroutines) that all read from in.
	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)
	c10 := factorial(in)

	fmt.Printf("During the process | Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("During the process | Number of Go routines: %d\n", runtime.NumGoroutine())
	//FAN IN
	// Multiplex multiple channels onto a single channel
	//merge channels from c0 to c9 onto a single channel
	var y int
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10) {
		y++
		fmt.Println(y, "\t", n)
	}
	fmt.Printf("End Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("End Number of Go routines: %d\n", runtime.NumGoroutine())
}

func inputGen() <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 100000; i++ {
			for j := 30; j < 40; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {

	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}

	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, n := range cs {
		go output(n)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
