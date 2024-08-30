package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

func rateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	limiter := rate.NewLimiter(2, 4)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			message := Message{
				Status: "request denied",
				Body:   "reached api capacity, try some time later",
			}

			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		}

		next(w, r)
	})

}
