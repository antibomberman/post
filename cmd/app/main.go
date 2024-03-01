package main

import (
	"antibomberman/post/internal/adapters/rest"
	"antibomberman/post/internal/config"
	"antibomberman/post/internal/di"
	"antibomberman/post/internal/repositories/postgres"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	cfg := config.Load()
	postgresDB := postgres.New(cfg)
	defer postgresDB.Close()
	err := postgresDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	container := di.NewDI(postgresDB, cfg)
	restServer := rest.New(container)

	log.Println("Server is running on port 8080")
	err = restServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
