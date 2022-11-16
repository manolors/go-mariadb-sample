package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/manolors/go-mariadb-sample/config"
	"github.com/manolors/go-mariadb-sample/models"

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

	db, err := gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{
		Logger: newLogger,
	})

	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	actor := models.Actor{}

	db.Debug().First(&actor)

	fmt.Printf("Primer actor: %s %s\n", actor.First_name, actor.Last_name)

	actores := []models.Actor{}
	result := db.Find(&actores)
	if result.Error == nil {
		fmt.Println("Registros encontrados: ", result.RowsAffected)
		for _, a := range actores {
			fmt.Printf("Actor: %s %s\n", a.First_name, a.Last_name)

		}
	}

}
