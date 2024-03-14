package main

import (
	_ "github.com/OddEer0/ck-filmoteka/docs"
	appRouter "github.com/OddEer0/ck-filmoteka/internal/presentation/router"
	"log"
	"net/http"
)

// @title VK-Filmoteka
// @version 1.0
// @description This is a sample HTTP package with Swagger annotations.
func main() {
	router := appRouter.NewAppRouter()
	initSwagger(router)

	server := http.Server{Addr: "localhost:5000", Handler: router}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server")
	}
}
