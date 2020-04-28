package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	out := make(chan int)
	go producer(ch, 100*time.Millisecond)
	go producer(ch, 250*time.Millisecond)
	go reader(out)

	for i := range ch {
		out <- i
	}
	time.Sleep(2 * time.Second)

}

func producer(ch chan int, d time.Duration) {
	var i int
	for {
		ch <- i
		i++
		time.Sleep(d)
	}
}

func reader(out chan int) {
	for n := range out {
		fmt.Println(n)
	}
}
