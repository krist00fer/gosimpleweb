package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server Started")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello World")
}
