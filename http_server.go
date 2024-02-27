package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Printf("hello called")
	fmt.Fprintf(w, "<p>hello there</p>")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8090", nil)

}
