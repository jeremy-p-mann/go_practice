package main

type Groceries struct {
	Month string   `json:"month"`
	Items []string `json:"fruits"`
}

func increment(x *int)  {
    *x = *x + 1
}

func main() {
    x := 1
    p := &x
    increment(p)
    println(x)

}
