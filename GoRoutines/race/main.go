package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Main")
	var counter int64 = 0
	maxRoutines := 100
	var wg sync.WaitGroup
	wg.Add(maxRoutines)
	for i := 0; i < maxRoutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Finished")
}
