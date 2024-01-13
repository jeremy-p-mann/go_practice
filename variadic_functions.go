package main

import "fmt"


func sum(nums ...int) {
    fmt.Println(nums)
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}


func main()  {
    sum()
    sum(2, 3, 4, 10)
    nums := []int{3, 4, 5}
    sum(nums...)
}

