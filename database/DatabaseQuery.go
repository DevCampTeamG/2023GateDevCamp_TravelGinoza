package database

import (
	"github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = DBConnect()
}
func UserStampRallyUpdate(uid string, sid uint) error {
	us := model.UserStampRally{UserId: uid, StampId: sid}

	result := db.Create(&us).Error

	return result
}

func UserStampRallyProgress(uid string) ([]model.UserStampRally, error) {
	us := []model.UserStampRally{}
	err := db.Find(&us).Where("user_id = ?", uid).Error

	return us, err

}
