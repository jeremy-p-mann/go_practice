package main

import "os"

func main() {
	// panic("A problem")
	_, err := os.Create("/tmp/defer.txt")
	if err != nil {
		panic(err)
	}

}
