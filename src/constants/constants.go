package constants

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DatabaseHost     string
	DatabaseUser     string
	DatabaseName     string
	DatabasePassword string
	DatabasePort     int
)

func init() {
	godotenv.Load()
	godotenv.Load("../.env")
	DatabaseHost = os.Getenv("DATABASE_HOST")
	DatabaseUser = os.Getenv("DATABASE_USER")
	DatabaseName = os.Getenv("DATABASE_NAME")
	DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	DatabasePort, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))
}
