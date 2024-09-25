package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		type Response struct {
			Message string `json:"message"`
		}
		resp := Response{
			Message: "Hello this is a hard corded json response!",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	log.Fatal(http.ListenAndServe(":8090", r))
}
