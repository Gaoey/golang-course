package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	total := 10

	wg.Add(total)
	start := time.Now()
	for i := 0; i < total; i++ {
		go printout(i)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}

func printout(i int) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println(i)
}
