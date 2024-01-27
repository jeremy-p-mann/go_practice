package main

import (
	"fmt"
	// "time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// go func() {
	// 	messages <- "hi"
	// }()
    // time.Sleep(time.Second)

	select {
	case msg := <-messages:
		fmt.Println("received", msg)
	default:
		fmt.Println("no message recieved")
	}
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no message received")
	}
}
