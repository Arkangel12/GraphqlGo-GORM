package db

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/Arkangel12/curso/graphql-gorm/config"
	"github.com/Arkangel12/curso/graphql-gorm/models"
)

var (
	once sync.Once
	db   *gorm.DB
	err  error
)

// GetConnection connects to database
func GetConnection() *gorm.DB {

	once.Do(func() {
		db, err = gorm.Open("postgres", "host="+config.AppConfig.Host+" port="+config.AppConfig.Port+" dbname="+config.AppConfig.Dbname+" user="+config.AppConfig.User+" password="+config.AppConfig.Password+" sslmode=disable")

		if err != nil {
			log.Fatal("Could not connecto to database ", err)
		}

		db.AutoMigrate(&models.User{}, &models.Article{})
	})

	return db
}
