package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}

func main() {
	//port := os.Getenv("PORT")
	http.HandleFunc("/", handler)
	//http.ListenAndServe(":"+port, nil)
	http.ListenAndServe(":8080", nil)
}
