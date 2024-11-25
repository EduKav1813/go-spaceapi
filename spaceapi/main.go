package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var jsonData interface{}

func main() {

	filePath := "../spaceinfo.json"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	if err := json.Unmarshal(byteValue, &jsonData); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	http.HandleFunc("/get", getHandler)
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(jsonData)
	if err != nil {
		http.Error(w, "Error preparing JSON response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
