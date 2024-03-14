package main

import (
	_ "github.com/OddEer0/ck-filmoteka/docs"
	"github.com/OddEer0/ck-filmoteka/internal/presentation/router"
	"log"
	"net/http"
)

// @title VK-Filmoteka
// @version 1.0
// @description This is a sample HTTP package with Swagger annotations.
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", router.NewAppRouter)

	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui/"))))
	mux.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	server := http.Server{Addr: "localhost:5000", Handler: mux}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server")
	}
}
