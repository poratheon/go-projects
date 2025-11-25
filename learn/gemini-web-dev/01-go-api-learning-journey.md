# Go API Learning Journey

## Module 1: Your First Go Web Server

**Goal:** Create a simple web server that responds with "Hello, World!". This will be the foundation of our API.

### 1. Understand the Code:

Go's standard library has a powerful `net/http` package that makes it easy to build web servers. Here’s how the pieces fit together:

*   **`package main`**: Declares that this is an executable program.
*   **`import (...)`**: Brings in necessary libraries. We'll need `fmt` for formatting text, `log` for error handling, and `net/http` for the web server functionality.
*   **`main()` function**: This is the entry point of our application.
*   **`http.HandleFunc("/", ...)`**: This function tells the server which function should handle incoming web requests to a specific path. In our case, all requests to the root path (`/`) will be handled by our `helloHandler`.
*   **`handler function`**: This function takes two arguments: an `http.ResponseWriter` (which is used to write the response back to the client) and an `*http.Request` (which contains information about the incoming request).
*   **`fmt.Fprintf`**: This is used to write the "Hello, World!" string to the `http.ResponseWriter`, sending it back to the client.
*   **`http.ListenAndServe(":8080", nil)`**: This starts the web server on port 8080. It will block and listen for incoming requests until the program is stopped. If it fails to start, it returns an error which we log using `log.Fatal`.

### 2. Update `main.go`:

Copy the following code and use it to replace the contents of your `main.go` file.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler) // Direct all traffic to helloHandler
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

### 3. Run and Test Your Server:

Once you've saved the `main.go` file with the above content, open your terminal and run the server:

```bash
go run main.go
```

You should see the message "Server starting on port 8080...". Now, open a *new* terminal window and use `curl` to test it:

```bash
curl http://localhost:8080
```

You should see the output: `Hello, World!`

---
