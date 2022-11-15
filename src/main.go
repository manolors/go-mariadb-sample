package main

import (
	"log"
	"os"
	"time"

	"github.com/manolors/gorm-crud/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	_, err := gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
}
