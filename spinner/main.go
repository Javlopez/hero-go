package main

import (
	"fmt"
	"time"
)

func main() {

	go spinner(100 * time.Millisecond)
	const number = 45
	fibN := fib(number)
	fmt.Printf("\rFibonacci(%d)=%d\n", number, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|⁄` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return 1
	}

	return fib(x-1) + fib(x-2)
}
