package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop() // Clean up when done, or if you break the loop.

	for {
		select {
		case t := <-ticker.C: // Read from the ticker channel.
			fmt.Println("Current time:", t)
		}
	}
}
