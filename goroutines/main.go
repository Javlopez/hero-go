package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("Hello from one")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello form two")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("This is the exit")
}
