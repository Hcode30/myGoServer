package main

import (
	"fmt"
	"log"
	"net/http"
)

func startServer(handler http.Handler, port string) {
	PORT := fmt.Sprintf(":%s", port)
	server := &http.Server{
		Handler: handler,
		Addr:    PORT,
	}
	log.Println("Server Started At Port:", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error Starting Server")
	}
}
