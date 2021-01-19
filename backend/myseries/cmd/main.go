package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	"github.com/mqrc81/myseries/postgres"
	"github.com/mqrc81/myseries/web"
)

func main() {
	port := ":3000"
	fmt.Println("Starting application...")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	store, err := postgres.NewStore(dataSourceName)

	if err != nil {
		log.Fatal(err)
	}
	chi.NewMux()

	// sessions, err := web.NewSessionManager(dataSourceName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	handler := web.NewHandler(store)

	fmt.Println("Listening on " + port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal(err)
	}
}
