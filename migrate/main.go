package main

import "github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/database"

func main() {
	db := database.DBConnect()
	database.DBMigrate(db)
}
