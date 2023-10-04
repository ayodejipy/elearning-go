package config

import "os"

type DatabaseConfig struct {
	Host string
	Port string
	DBName string
	User string
	Password string
}


func LoadConfig() *DatabaseConfig {
	// configuration
	config := &DatabaseConfig{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		User: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	// return the credentials
	return config
}