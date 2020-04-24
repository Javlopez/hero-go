package main

import "fmt"

func main() {

	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10000; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for n := range ch {
			fmt.Println(n)
		}
		done <- true
	}()

	go func() {
		for n := range ch {
			fmt.Println(n)
		}
		done <- true
	}()

	<-done
	<-done
}
