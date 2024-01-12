package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error)  {
    if e != nil {
        panic(e)
    }
}

func main()  {
    filename := "/tmp/EpocCamServiceLog.txt"
    dat, err := os.ReadFile(filename)
    check(err)
    fmt.Println(string(dat[:3]))
    fmt.Println(dat[:3])
    f, err := os.Open(filename)
    defer f.Close()

    b1 := make([]byte, 5)
    _, err = f.Read(b1)
    println(string(b1))

    c1 := make([]byte, 5)
    _, err = f.Seek(1, 0)
    check(err)
    _, err = f.Read(c1)
    println(string(c1))

    _, err = f.Seek(0, 0)
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    println(string(b4))
}
