package main

import "fmt"


func main()  {
    messages := make(chan string, 1)

    messages <- "foo"
    messages <- "bar"

    fmt.Println(<-messages)
    fmt.Println(<-messages)

}
