package main

import (
	"fmt"
)

func main() {
	queue := inputGen(10, 10)
	factory := factorialGen(queue)
	for res := range factory {
		fmt.Println(res)
	}

}

func inputGen(mFirst, mSecond int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < mFirst; i++ {
			for j := 0; j < mSecond; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorialGen(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- factorial(n)
		}
		close(out)
	}()
	return out
}

func factorial(n int) int {

	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}

	return total

}
