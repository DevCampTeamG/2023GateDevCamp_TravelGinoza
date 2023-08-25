package model

import "gorm.io/gorm"

type UserStampRally struct {
	gorm.Model `json:"-"`
	UserId     string `gorm:"uniqueIndex:unique_user_stamp_rally;not null"`
	StampId    uint   `gorm:"uniqueIndex:unique_user_stamp_rally;not null"`
}

type UserStampRallyProgress struct {
	UserId string
	Stamp1 bool
	Stamp2 bool
	Stamp3 bool
	Stamp4 bool
}
