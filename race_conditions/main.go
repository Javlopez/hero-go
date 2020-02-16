package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	incr := 0
	gr := 100
	wg.Add(gr)
	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			m.Lock()
			v := incr
			v++
			incr = v
			fmt.Printf("Incrementor:  %d\n", incr)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("\nEnd Value:  %d\n\n", incr)

}
