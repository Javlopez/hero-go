package main

import (
	"fmt"

	"time"

	"math/rand"
)

func main() {
	c := fanIn(boring("Jhon"), boring("Jav"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're both boring; I'm leaving...")
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s-%d", msg, i)
			time.Sleep(time.Duration(rand.Int31n(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
