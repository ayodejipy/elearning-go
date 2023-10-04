package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ayodejipy/elearning-go/internal/config"
	"github.com/ayodejipy/elearning-go/internal/database"
	"github.com/ayodejipy/elearning-go/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	router := routes.NewRouter()

	// load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("unable to load env file: %v", err)
	}
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	// attempt to connect to database
	config := config.LoadConfig()
	database.InitDatabaseConnection(*config)


	// run server
	fmt.Printf("Server started and running on http://localhost/%s \n", addr);

	log.Fatal(http.ListenAndServe(addr, router))
}