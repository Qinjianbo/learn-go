package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello goroutine")
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")
}
