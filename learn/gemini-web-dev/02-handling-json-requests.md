## Module 2: Handling JSON Requests

**Goal:** Turn our simple web service into a real API by teaching it to handle JSON. Our API will accept a JSON object with a question and return a JSON object with a dummy answer.

### 1. Understand the Concepts:

*   **Structs for JSON:** In Go, we define the "shape" of JSON objects using `structs`. We use `json` tags (the part in backticks) to tell Go how to map the fields in our struct to the keys in the JSON object.
*   **Decoding (Request):** To read a JSON request, we use `json.NewDecoder(r.Body).Decode(&ourStruct)`. This is a highly efficient, streaming decoder that reads from the request body (`r.Body`) and populates our struct (`&ourStruct`).
*   **Encoding (Response):** To send a JSON response, we first set the correct HTTP header: `w.Header().Set("Content-Type", "application/json")`. Then, we use `json.NewEncoder(w).Encode(ourData)` to write our Go struct or map as a JSON string directly to the `ResponseWriter`.

### 2. Update `main.go`:

Replace the code in `main.go` with the following. This version defines the request/response structs and implements a `/api/search` endpoint that processes JSON.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// SearchRequest defines the structure of the incoming JSON request.
// The `json:"..."` tags map the struct fields to JSON keys.
type SearchRequest struct {
	Question string `json:"question"`
	Limit    int    `json:"limit"`
}

// SearchResponse defines the structure of the JSON response.
type SearchResponse struct {
	Answer   string   `json:"answer"`
	Sources  []string `json:"sources"`
	Duration string   `json:"duration_ms"`
}

// searchHandler handles requests to the /api/search endpoint.
func searchHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// 1. Ensure the request is a POST request.
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Decode the incoming JSON payload into our SearchRequest struct.
	var req SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// We'll log the received question to the server console.
	log.Printf("Received question: '%s' with limit %d", req.Question, req.Limit)

	// 3. For now, create a dummy response. In the future, this data
	// will come from our Snowflake RAG logic.
	response := SearchResponse{
		Answer:   fmt.Sprintf("This is a dummy answer to '%s'", req.Question),
		Sources:  []string{"snowflake-table-1", "snowflake-table-2"},
		Duration: fmt.Sprintf("%d", time.Since(start).Milliseconds()),
	}

	// 4. Encode the response struct as JSON and send it back to the client.
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func main() {
	// We now point the /api/search route to our new JSON handler.
	http.HandleFunc("/api/search", searchHandler)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

### 3. Run and Test Your API:

1.  **Run the server:**
    ```bash
    go run main.go
    ```

2.  **Test with `curl`:** Open a new terminal. This `curl` command is more complex.
    *   `-X POST`: Specifies the HTTP POST method.
    *   `-H "Content-Type: application/json"`: Sets the content type header.
    *   `-d '{"question": "what is the meaning of life?", "limit": 2}'`: Provides the JSON request body.

    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"question": "what is the meaning of life?", "limit": 2}' http://localhost:8080/api/search
    ```

    You should see a JSON response similar to this:
    ```json
    {"answer":"This is a dummy answer to 'what is the meaning of life?'","sources":["snowflake-table-1","snowflake-table-2"],"duration_ms":"0"}
    ```
    *(Note: The duration will likely be 0 or a very small number).*

### 4. (Bonus) Testing from the Browser Console

You cannot test a `POST` endpoint by simply typing a URL into the browser's address bar, but you *can* use the browser's built-in **Developer Tools**. This is how a frontend web application would communicate with your API.

1.  Make sure your Go application is running.
2.  Open a new, empty tab in your browser (typing `about:blank` works well).
3.  Open the Developer Tools (usually `F12` or `Cmd+Opt+I` on Mac) and click the **"Console"** tab.
4.  Paste this JavaScript `fetch` command into the console and press Enter:

```javascript
fetch('http://localhost:8080/api/search', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    question: "what is the meaning of life? (from browser)",
    limit: 3
  }),
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
```
You will see the JSON response from your API logged directly in the console.

---