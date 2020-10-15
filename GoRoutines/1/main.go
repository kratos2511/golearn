package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("GOOS", runtime.GOOS)
	fmt.Println("GOARCH", runtime.GOARCH)
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("NumGoroutines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	maxRoutines := 5
	wg.Add(maxRoutines)
	for i := 0; i < maxRoutines; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Main")

}
