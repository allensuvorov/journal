package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Journal"))
}

func main() {
	fmt.Println("Hello world!")
}
