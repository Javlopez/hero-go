package main

import "fmt"

func main() {

	c := make(chan int)

	go send(c)

	receive(c)

	fmt.Println("Finishing the execution")

}

//send
func send(c chan<- int) {
	c <- 42
}

//receive
func receive(c <-chan int) {
	fmt.Printf("Receiving daata from channel c: %d\n", <-c)
}
