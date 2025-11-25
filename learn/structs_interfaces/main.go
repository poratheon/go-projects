package main

import (
	"encoding/json"
	"fmt"
)

// structs and interfaces - refresher

/*
ETL structs are usually used in two ways:
1) DTO - Data transfer objects - define the shape of the data (json) both coming in or going out
2) Configuration - used to hold state (e.g. db connection)
*/

// Basic definition and tags
// in web dev - "tags" which are the text in backticks are critical - they provide the mapping between Go struct fields to JSON and/or Database

// Cortex request - a request from a user to the API
type CortexRequest struct {
	// tag is query_text - so in JSON it's the key where the value can be extracted from and mapped here
	QueryText string `json:"query_text"`
	// this tag has the "omitempty" so if it's zero, it's ignored
	Limit     int `json:"limit,omitempty"`
	processed bool
}

func (r *CortexRequest) Validate() error {
	if r.QueryText == "" {
		return fmt.Errorf("query_text cannot be empty")
	}
	if r.Limit > 100 {
		return fmt.Errorf("limit cannot exceed 100")
	}
	return nil
}

func ParseAndValidateRequest(data []byte) (*CortexRequest, error) {
	var req CortexRequest

	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("bad json: %w", err)
	}

	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	return &req, nil
}

func main() {
	rawJSON := []byte(`{"query_text":"Does Ford own any licenses of NX?", "limit": 5}`)

	req, err := ParseAndValidateRequest(rawJSON)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Parsed Struct:  %+v\n", *req)
	fmt.Println("Question:", req.QueryText)

	rawJSON2 := []byte(`{"query_text":"Does this show the limit?", "limit": 0}`)

	var req2 *CortexRequest
	req2, err = ParseAndValidateRequest(rawJSON2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Parsed second struct:  %+v\n", *req2)
	fmt.Println("Question:", req2.QueryText)

}
