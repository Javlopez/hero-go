package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("Start Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("Start Number of Go routines: %d\n", runtime.NumGoroutine())
	for i := 0; i < 24; i++ {
		fmt.Printf("During the process y| Number of CPUS: %d\n", runtime.NumCPU())
		fmt.Printf("During the process y| Number of Go routines: %d\n", runtime.NumGoroutine())
		c := timer(i)
		fmt.Println(<-c)
	}
	fmt.Printf("End Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("End Number of Go routines: %d\n", runtime.NumGoroutine())
}

func timer(n int) <-chan int {

	ch := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Printf("During the process x| Number of CPUS: %d\n", runtime.NumCPU())
		fmt.Printf("During the process x| Number of Go routines: %d\n", runtime.NumGoroutine())
		ch <- n
	}()
	return ch
}
