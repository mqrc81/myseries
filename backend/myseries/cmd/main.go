package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/mqrc81/myseries/backend/postgres"
	"github.com/mqrc81/myseries/backend/web"
)

func main() {
	fmt.Println("Starting application...")

	if os.Getenv("ENVIRONMENT") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
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

	sessions, err := web.NewSessionManager(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	csrfKey := make([]byte, 32)
	if _, err = rand.Read(csrfKey); err != nil {
		log.Fatalf("error generating csrf-key: %v", err)
	}

	handler := web.NewHandler(store, sessions, csrfKey)

	fmt.Println("Listening on port " + os.Getenv("PORT"))
	if err = http.ListenAndServe(":"+os.Getenv("PORT"), handler); err != nil {
		log.Fatal(err)
	}
}
