package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/didip/tollbooth/v7"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {

	message := &Message{
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
	message := &Message{
		Status: "request denied",
		Body:   "reached api capacity, try some time later",
	}

	jsonMessage, _ := json.Marshal(message)

	// setting 2 requests per second
	tlbthLimiter := tollbooth.NewLimiter(2, nil)

	// setting burst to 4
	tlbthLimiter.SetBurst(4)

	tlbthLimiter.SetMessageContentType("application/json")
	tlbthLimiter.SetMessage(string(jsonMessage))

	http.Handle("/ping", tollbooth.LimitFuncHandler(tlbthLimiter, endpointHandler))

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
