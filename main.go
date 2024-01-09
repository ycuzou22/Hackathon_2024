package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	startServ()
}

func startServ() {
	fmt.Println("http://localhost" + port + " 🚀")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello World") })
	http.ListenAndServe(port, nil)
}
