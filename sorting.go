package main

import (
    "fmt"
    "slices"
)


func main()  {
    strs := []string{"c", "a", "b"}
    slices.Sort(strs)
    fmt.Println("strings:", strs)

    ints := []int{7, 2, 4}
    fmt.Println("ints:", ints)
    slices.Sort(ints)
    fmt.Println("ints:", ints)

    s := slices.IsSorted(ints)
    fmt.Println("Sorted:", s)
}
