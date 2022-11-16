package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func main() {
	godotenv.Load("../.env")
	config := DBConfig{
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("ROOT_PASSWORD"),
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		DBName:   os.Getenv("DATABASE_NAME"),
	}
	fmt.Printf("dbconfig: %#v\n", config)
	db, err := sql.Open("mysql", DbURL(&config))
	if err != nil {
		log.Println("Error connecting database")
	}
	defer db.Close()

	results, err := db.Query("SELECT first_name, last_name FROM actor")
	if err != nil {
		log.Println("Error retrieving data")
	}
	type Actor struct {
		first_name string
		last_name  string
	}
	for results.Next() {
		var actor Actor
		err = results.Scan(&actor.first_name, &actor.last_name)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("Actor: %s %s\n", actor.first_name, actor.last_name)
	}

}
