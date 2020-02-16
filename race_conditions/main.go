package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	incr := 0
	gr := 100

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			v := incr
			runtime.Gosched()
			v++
			incr = v
			fmt.Printf("Incrementor:  %d\n", incr)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("\nIncrementor:  %d\n\n", incr)

}
