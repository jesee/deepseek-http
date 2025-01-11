package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	fmt.Println("Starting server on :8089")
	if err := http.ListenAndServe(":8089", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}