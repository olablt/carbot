package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type TargetFunc[In any, Out any] func(context.Context, In) (Out, error)

func Handle[In any, Out any](f TargetFunc[In, Out]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var in In

		// Retrieve data from request.
		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			// Format error response
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		// Call out to target function
		out, err := f(r.Context(), in)
		if err != nil {
			// Format error response
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Format and write response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			log.Printf("failed to encode created note: %v", err)
			return
		}
	})
}
