package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Hello, World!")

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT environment variable was not set")
	}

	fmt.Println("PORT is set to ", portString)
}
