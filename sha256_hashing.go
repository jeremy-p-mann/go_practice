package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "this is a sha string"
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Println("%x\n", bs)
}
