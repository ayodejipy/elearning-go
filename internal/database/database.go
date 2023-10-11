package database

import (
	"fmt"
	"log"

	"github.com/ayodejipy/elearning-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConn struct {
	*gorm.DB
}

func InitDatabaseConnection (config config.DatabaseConfig) (*DBConn)  {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Password, config.DBName, config.Port)

	// initialize connection to database using gorm
	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	// catch("Unable to connect to database.", err)
	if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
		return nil
    }


	fmt.Println("Database connection established.")
	// return database
	return &DBConn{database}
}


