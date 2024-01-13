package main

import "fmt"


type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num %v", b.num)
}

type container struct {
    base
    str string
}

func main()  {

    c := container {
        base: base{
            num: 1,
        },
        str: "foo",
    }
    fmt.Println(c.num, c.str)
    fmt.Println(c.base.num)
    fmt.Println(c.describe())
    type describer interface {
        describe() string
    }
    var d describer = c
    fmt.Println("describe: ", d.describe())
}
