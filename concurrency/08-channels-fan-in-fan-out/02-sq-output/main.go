package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	//FAN OUT
	//Distribute the sq work across two goroutines that both read from in
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// Consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	output := make(chan int)
	go func() {
		for _, n := range nums {
			output <- n
		}
		close(output)
	}()
	return output
}

func sq(c chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	fmt.Printf("TYPE OF CS: %T\n", cs)
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}
