package main

import (
	"fmt"
	"strconv"
)


func main()  {
    f, _ := strconv.ParseFloat("3.14145", 64)
    fmt.Println(f)
    i, _ := strconv.ParseInt("3", 0, 64)
    fmt.Println(i)
    s, _ := strconv.ParseUint("-3", 0, 64)
    fmt.Println(s)
}
