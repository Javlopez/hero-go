package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Start Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("Start Number of Go routines: %d\n", runtime.NumGoroutine())

	c := Incrementor(2)

	var count int
	fmt.Printf("During the process | Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("During the process | Number of Go routines: %d\n", runtime.NumGoroutine())
	for n := range c {
		count++
		fmt.Println(n)
	}
	fmt.Println("Final Counter:", count)
	fmt.Printf("End Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("End Number of Go routines: %d\n", runtime.NumGoroutine())
}

func Incrementor(n int) <-chan string {
	out := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for j := 0; j < 20; j++ {
				out <- fmt.Sprintf("Process %d printing %d", i, j)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(out)
	}()

	return out
}
