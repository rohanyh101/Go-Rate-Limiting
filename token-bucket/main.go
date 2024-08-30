package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {

	message := Message{
		Status: "success",
		Body:   "HI, you have reached the endpoint!, how may I help you ?",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/ping", rateLimiter(endpointHandler))

	log.Println("server is running on port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
	}
}
