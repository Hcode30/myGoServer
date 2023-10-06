package main

import (
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is Not Available")
	}

	r := chi.NewRouter()
	startServer(r, port)

}
