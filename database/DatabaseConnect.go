package database

import (
	"github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../teamG"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.UserStampRally{})
}
