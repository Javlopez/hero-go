package main

import "fmt"

func main() {
	//f := factorial(4)
	fmt.Println("Total:", simpleFactorial(44))
	f := factorial(44)
	for n := range f {
		fmt.Println(n)
	}
	//fmt.Println(goFact())
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()

	return out
}

func simpleFactorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}

	return total
}
