package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Printf("The number stored in the channel c is: %d\n", <-c)
}
