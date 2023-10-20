package crud

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type CrudRequest struct {
	gorm.Model
	Action string `json:"action"`
	Fields Fields
}

type Fields struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Column []Column
}

type Column struct {
	Name  string
	Value any
}

func main() {
	// Create a new ServeMux (router)
	mux := http.NewServeMux()
	// Define a route "/test/go" and associate it with a handler function
	mux.HandleFunc("/crud/go", GoCrud)

	// Create an HTTP server with the chosen port and the ServeMux
	port := 8080
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	// Start the HTTP server
	fmt.Printf("Server listening on port %d...\n", port)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func GoCrud(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed for this endpoint", http.StatusBadRequest)
	}
	var requestData CrudRequest
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Failed to decode JSON request", http.StatusBadRequest)
	}

	Action(w, &requestData)
}

func Action(w http.ResponseWriter, data *CrudRequest) {

	switch strings.ToLower(data.Action) {
	case "create":
		// run create function
	case "read":
		// run read function
	case "update":
		// run update function
	case "delete":
		// run delete function
	default:
		http.Error(w, "invalid action, available actions: CRUD", http.StatusNotFound)
	}
}

func CrudCreate(data *CrudRequest) {

}
