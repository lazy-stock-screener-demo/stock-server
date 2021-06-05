package gormclient

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Client is singleton of db client/pool
var ClientGorm *gorm.DB

// NewConnectedGorm define a connection to DB
func NewConnectedGorm() *gorm.DB {
	if ClientGorm == nil {
		config := NewDBConfig()

		database, err := gorm.Open(
			postgres.Open(fmt.Sprintf(
				"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
				config.Addr,
				config.Port,
				config.Username,
				config.DBName,
				config.Password,
			)), &gorm.Config{},
		)
		if err != nil {
			fmt.Print(fmt.Errorf("Connecting to DB Error: %s", err))
		}
		db, _ := database.DB()
		fmt.Println("Connected to PostgreSQL!")
		db.SetMaxOpenConns(200)
		ClientGorm = database
		return database
	}
	return ClientGorm
}
