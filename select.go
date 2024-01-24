package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	c3 := make(chan string)
	c4 := make(chan string)
    chans := []chan string{c3, c4}

	go func() {
		time.Sleep(time.Second)
		c3 <- "three"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "four"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chans[0]:
			fmt.Println("received", msg1)
		case msg2 := <-chans[1]:
			fmt.Println("received", msg2)
		}
	}


}
