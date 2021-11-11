package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8090", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("This is a simple response")
}
