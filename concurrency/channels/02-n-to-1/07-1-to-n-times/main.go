package main

import "fmt"

func main() {

	n := 10
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10000; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for n := range ch {
				fmt.Println(n)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}
}
