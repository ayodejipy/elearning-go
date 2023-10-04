package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	// run server
	fmt.Printf("Server started and running on http://localhost/%s \n", addr);

	log.Fatal(http.ListenAndServe(addr, router))
}