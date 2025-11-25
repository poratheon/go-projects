package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	http.HandleFunc("/", helloHandler)
	portNum := 8080
	port := ":" + strconv.Itoa(portNum)
	fmt.Printf("Starting server on port %v\n", portNum)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
