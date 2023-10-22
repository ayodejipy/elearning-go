package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ayodejipy/elearning-go/internal/config"
	"github.com/ayodejipy/elearning-go/internal/database"
	"github.com/ayodejipy/elearning-go/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("unable to load env file: %v", err)
	}

	// load config and init a connection to database
	config := config.LoadConfig()
	db := database.InitDatabaseConnection(*config)
	database.SyncDatabase(db)


	// init a new chi router instance
	r := chi.NewRouter()
	routes.SetupRoutes(r, db)

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	// run server
	fmt.Printf("Server started and running on http://localhost%s \n", addr);

	log.Fatal(http.ListenAndServe(addr, r))
}