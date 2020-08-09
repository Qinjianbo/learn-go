package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	fmt.Println("producer")
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if (ok == false) {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	go producer(ch)
	time.Sleep(100 * time.Millisecond)
	// 使用range
	for v := range ch {
		fmt.Println("Received1 ", v)
	}
}
