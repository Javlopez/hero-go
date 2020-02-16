package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 42
	c <- 43

	fmt.Printf("The number stored in the channel c is: %d\n", <-c)
	fmt.Printf("The number stored in the channel c is: %d\n", <-c)
}
