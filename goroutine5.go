package main

import (
	"fmt"
)

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	sendch := make(chan<- int)
	go sendData(sendch)
	fmt.Println(<-sendch)
	// 运行后报错
	// invalid operation: <-sendch (receive from send-only type chan<- int)
}
