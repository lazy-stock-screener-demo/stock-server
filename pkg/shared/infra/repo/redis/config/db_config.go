package redisclient

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

// DBConfig define a struct of configuration
type DBConfig struct {
	Addr     string
	Port     string
	Username string
	DBName   int
	Password string
}

// NewDBConfig define Configuration setting of db
func NewDBConfig() *DBConfig {
	Addr := os.Getenv("IDENTITY_DEV_HOST")
	Port := os.Getenv("IDENTITY_DEV_PORT")
	Username := os.Getenv("IDENTITY_DEV_USER")
	Password := os.Getenv("IDENTITY_DEV_PASSWORD")
	DBName := os.Getenv("IDENTITY_DEV_DB_NAME")
	db, _ := strconv.Atoi(DBName)
	return &DBConfig{
		Addr:     Addr,
		Port:     Port,
		Username: Username,
		Password: Password,
		DBName:   db,
	}
}
