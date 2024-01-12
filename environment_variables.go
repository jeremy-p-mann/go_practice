package main

import (
	"os"
	"strings"
)


func main()  {
    os.Setenv("FOO", "fadsf")
    println(os.Getenv("FOO"))

    for _, e := range os.Environ() {
        println(e)
        println(strings.Split(e, "=")[0])
        println(strings.Split(e, "=")[1])
    }
}
