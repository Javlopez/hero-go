package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		wg   sync.WaitGroup
		incr int64
	)
	gr := 100
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&incr, 1)
			fmt.Println(atomic.LoadInt64(&incr))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("\nEnd Value:  %d\n\n", incr)
}
