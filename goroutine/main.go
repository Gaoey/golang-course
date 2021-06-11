package main

import (
	"fmt"
	"time"
)

func main() {
	total := 10

	channelSignal := make(chan struct{})

	start := time.Now()
	for i := 0; i < total; i++ {
		go printout(i, channelSignal)
	}

	for i := 0; i < total; i++ {
		<-channelSignal
	}

	fmt.Println(time.Since(start))
}

// มี signal ไปบอกข้างนอกว่าเสร็จแล้ว
func printout(i int, channelSignal chan struct{}) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println(i)
	channelSignal <- struct{}{}
}
