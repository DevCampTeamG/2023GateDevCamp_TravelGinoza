package database

import (
	"github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../teamG"), &gorm.Config{})
	db.Logger = db.Logger.LogMode(logger.Info)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.UserStampRally{})
}
