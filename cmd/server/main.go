package main

import (
	"log"
	"net/http"

	"github.com/chin13577/palmon-go-server-test/internal/httpx"
)

func health(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"status": "ok"}
	httpx.OK(w, data, "server is healthy")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /boom", func(w http.ResponseWriter, r *http.Request) {
		httpx.Fail(w, httpx.CodePlayerNotFound, "Player not found")
	})

	log.Println("listening on http://localhost:4001")
	err := http.ListenAndServe(":4001", mux)

	if err != nil {
		log.Fatal(err)
	}
}
