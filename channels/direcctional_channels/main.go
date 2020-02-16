package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("------------")
	fmt.Printf("%T\n", c)

	cr := make(<-chan int, 2) //type channel int: receiveing from a channle
	cs := make(chan<- int, 2) //type channel int: sending to channel

	fmt.Println("----TYPES --------")
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// general to specific will work
	// specific to general does not work
	fmt.Println("----CONVERT --------")
	fmt.Printf("%T\n", (chan<- int)(c))
	fmt.Printf("%T\n", (<-chan int)(c))
}
