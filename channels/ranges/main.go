package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Everyting is done")
}

func send(c chan<- int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)
}
