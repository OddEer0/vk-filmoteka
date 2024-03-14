package main

import (
	_ "github.com/OddEer0/ck-filmoteka/docs"
	"github.com/OddEer0/ck-filmoteka/internal/presentation/router"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	server := http.Server{Addr: "localhost:5000", Handler: &router.AppRouter{}}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server")
	}
}
