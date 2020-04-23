package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter int64
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
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(label, i, "Counter:", counter)
	}

	wg.Done()
}
