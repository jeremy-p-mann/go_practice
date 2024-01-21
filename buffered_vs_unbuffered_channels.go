package main

import (
	"log"
	"time"
)

func main() {
	bufferedChan := make(chan string, 1)

    i := "I"
	go func() {
		log.Println("I am sending message to buffered channel...")
		bufferedChan <- "Hello from buffered channel" + i
		log.Println("I am free")
        i="kdsjf"
	}()

	log.Println("Receiving message from buffered channel...")
	msg := <-bufferedChan
	log.Println("Received:", msg)

	log.Printf("%T", bufferedChan)
	log.Println("Received:", msg)
	time.Sleep(1 * time.Second) // To observe the output

	unbufferedChan := make(chan string)
	log.Println("------------------")

	go func() {
		log.Println("I am sending message to unbuffered channel...")
		unbufferedChan <- "Hello from unbuffered channel"
		log.Println("I am free")
	}()

	log.Println("Receiving message from unbuffered channel...")
	msg = <-unbufferedChan
	log.Println("Received:", msg)
	time.Sleep(1 * time.Second) // To observe the output
}
