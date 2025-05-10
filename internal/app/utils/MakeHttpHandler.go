package utils

import (
	"log"
	"net/http"
)

func MakeHttpHandler(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("handler error: %v", err)
		}
	}
}
