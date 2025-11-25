package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type SearchRequest struct {
	Question string `json:"question"`
	Limit    int    `json:"limit"`
}

type SearchResponse struct {
	Answer   string   `json:"answer"`
	Sources  []string `json:"sources"`
	Duration string   `json:"duration_ms"`
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// check to ensure POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed.  Method must be POST.", http.StatusMethodNotAllowed)
		return
	}

	// decode the request JSON payload
	var req SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}
	// log question to console
	log.Printf("Received question:  '%s' with limit %d", req.Question, req.Limit)

	// dummy response based on the payload receved
	response := SearchResponse{
		Answer:   fmt.Sprintf("This is a dummy response to '%s'", req.Question),
		Sources:  []string{"snowflake-table-abcxyz", "snowflake-table-123789"},
		Duration: fmt.Sprintf("Took %d milliseconds", time.Since(start).Milliseconds()),
	}

	// encoding and returning the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("There was an error encoding the response:  %v", err)
		http.Error(w, "Failed to write a response", http.StatusInternalServerError)
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api/search", searchHandler)
	portNum := 8080
	port := ":" + strconv.Itoa(portNum)
	fmt.Printf("Starting server on port %v\n", portNum)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
