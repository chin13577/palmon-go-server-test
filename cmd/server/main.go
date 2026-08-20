package main

import (
	"log"
	"net/http"

	"github.com/chin13577/palmon-go-server-test/internal/config"
	"github.com/chin13577/palmon-go-server-test/internal/httpx"
	"github.com/joho/godotenv"
)

func health(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"status": "ok"}
	httpx.OK(w, data, "server is healthy")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found, using environment variables")
	}

	cfg := config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /boom", func(w http.ResponseWriter, r *http.Request) {
		httpx.Fail(w, httpx.CodePlayerNotFound, "Player not found")
	})

	addr := ":" + cfg.Port
	log.Printf("listening on http://localhost%s", addr)
	err = http.ListenAndServe(addr, mux)

	if err != nil {
		log.Fatal(err)
	}
}
