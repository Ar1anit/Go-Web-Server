package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT environment variable was not set")
	}

	fmt.Println("PORT is set to ", portString)

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr   : ":" + portString,
	}

	log.Printf("Server is running on port %s", portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
