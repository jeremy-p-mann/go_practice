package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 10000.
	const d = 3e20 / n
	fmt.Printf("variable count=%v is of type %T \n", d, d)
	fmt.Printf("variable count=%v is of type %T \n", n, n)
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(d))

}
