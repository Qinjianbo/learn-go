package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello goroutine!")
}

func main() {
	go hello()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("main function")
}
