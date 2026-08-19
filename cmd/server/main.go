package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chin13577/palmon-go-server-test/internal/httpx"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := httpx.APIResponse{
		StatusCode: 200,
		Message:    "server is healthy",
		Data:       map[string]string{"status": "ok"},
	}
	_ = json.NewEncoder(w).Encode(body)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)

	log.Println("listening on http://localhost:4001")
	err := http.ListenAndServe(":4001", mux)

	if err != nil {
		log.Fatal(err)
	}
}
