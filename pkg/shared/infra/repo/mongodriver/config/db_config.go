package mongoclient

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
	Addr := os.Getenv("CATALOG_DEV_HOST")
	Port := os.Getenv("CATALOG_DEV_PORT")
	Username := os.Getenv("CATALOG_DEV_USER")
	Password := os.Getenv("CATALOG_DEV_PASSWORD")
	DBName := os.Getenv("CATALOG_DEV_DB_NAME")
	return &DBConfig{
		Addr:     Addr,
		Port:     Port,
		Username: Username,
		Password: Password,
		DBName:   DBName,
	}
}
