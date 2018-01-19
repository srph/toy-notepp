package settings

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

var (
	// App settings
	AppVersion string
	AppName string
	AppDescription string

	// MySQL settings
	DatabaseHost string
	DatabaseUser string
	DatabasePassword string
	DatabaseName string
	DatabasePort string
)

func Init() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
	}

	AppVersion = "1.0"
	AppName = "Note++"
	AppDescription = "A toy note-taking web app"

	DatabaseHost = os.Getenv("DB_HOST")
	DatabaseUser = os.Getenv("DB_USER")
	DatabasePassword = os.Getenv("DB_PASSWORD")
	DatabaseName = os.Getenv("DB_NAME")
	DatabasePort = os.Getenv("DB_PORT")
}