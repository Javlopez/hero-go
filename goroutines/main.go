package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	fmt.Printf("Start Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("Start Number of Go routines: %d\n", runtime.NumGoroutine())

	fmt.Println()

	wg.Add(2)
	go func() {
		fmt.Println("Hello from one")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello form two")
		wg.Done()
	}()
	fmt.Printf("During the process | Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("During the process | Number of Go routines: %d\n", runtime.NumGoroutine())
	fmt.Println()
	wg.Wait()
	fmt.Println("This is the exit")

	fmt.Printf("End Number of CPUS: %d\n", runtime.NumCPU())
	fmt.Printf("End Number of Go routines: %d\n", runtime.NumGoroutine())
	fmt.Println()
}
