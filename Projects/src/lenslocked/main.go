package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Namaste  Parashar Ji!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", nil)
}
