package main

import (
	"encoding/json"
	"log"
	"net/http"
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

func searchHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	err := json.NewDecoder(r.Body). 
}
func main() {
	
}
