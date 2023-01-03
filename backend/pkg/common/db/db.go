package db

import (
	"log"

	"github.com/kayraberktuncer/sports-planner/pkg/common/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() models.Store {
	dbUrl := viper.Get("DB_URL").(string)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Exercise{})

	return db
}
