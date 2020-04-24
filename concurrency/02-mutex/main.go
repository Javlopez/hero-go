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
	mutex   sync.Mutex
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
		mutex.Lock()
		counter++
		fmt.Println(label, i, "Counter:", counter)
		mutex.Unlock()
	}

	wg.Done()
}
