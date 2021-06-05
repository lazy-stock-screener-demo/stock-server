package gormclient

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// DBConfig define a struct of configuration
type DBConfig struct {
	Addr     string
	Port     string
	Username string
	DBName   string
	Password string
}

// NewDBConfig define Configuration setting of db
func NewDBConfig() *DBConfig {
	Addr := os.Getenv("CUSTOMER_SELF_DEV_HOST")
	Port := os.Getenv("CUSTOMER_SELF_DEV_PORT")
	Username := os.Getenv("CUSTOMER_SELF_DEV_USER")
	Password := os.Getenv("CUSTOMER_SELF_DEV_PASSWORD")
	DBName := os.Getenv("CUSTOMER_SELF_DEV_DB_NAME")
	return &DBConfig{
		Addr:     Addr,
		Port:     Port,
		Username: Username,
		Password: Password,
		DBName:   DBName,
	}
}
