package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {

	wg.Add(2)
	go Incrementor("Foo:")
	go Incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func Incrementor(label string) {

	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		counter = x

		fmt.Println(label, i, "Counter:", counter)
	}

	wg.Done()
}
