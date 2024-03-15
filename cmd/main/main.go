package main

import (
	_ "github.com/OddEer0/ck-filmoteka/docs"
	"github.com/OddEer0/ck-filmoteka/internal/infrastructure/config"
	slogger "github.com/OddEer0/ck-filmoteka/internal/infrastructure/logger"
	appRouter "github.com/OddEer0/ck-filmoteka/internal/presentation/router"
	"log"
	"net/http"
)

// @title VK-Filmoteka
// @version 1.0
// @description This is a sample HTTP package with Swagger annotations.
func main() {
	cfg := config.MustLoad()
	logger := slogger.SetupLogger(cfg.Env)
	logger.Info("Logger setup")
	router := appRouter.NewAppRouter(logger)
	logger.Info("router setup")
	initSwagger(router)
	logger.Info("swagger setup")

	server := http.Server{Addr: cfg.Server.Address, Handler: router}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server")
	}
}

func initSwagger(mux *http.ServeMux) {
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui/"))))
	mux.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})
}
