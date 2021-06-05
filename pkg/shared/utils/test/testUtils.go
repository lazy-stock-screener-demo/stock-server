package testutils

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "/Users/universetennis/Code/dan-stock-screener-project-orchestration/services/stock"

// LoadEnv loads env vars from .env
func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		fmt.Println("Load Env failed", err)
		// log.WithFields(log.Fields{
		// 	"cause": err,
		// 	"cwd":   cwd,
		// }).Fatal("Problem loading .env file")

		os.Exit(-1)
	}
}
