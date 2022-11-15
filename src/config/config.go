// Config/Database.go
package config

import (
	"fmt"

	"github.com/manolors/go-mariadb-sample/constants"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     constants.DatabaseHost,
		Port:     constants.DatabasePort,
		User:     constants.DatabaseUser,
		Password: constants.DatabasePassword,
		DBName:   constants.DatabaseName,
	}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	a := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	fmt.Println("Conectando a", a)
	return a
}
