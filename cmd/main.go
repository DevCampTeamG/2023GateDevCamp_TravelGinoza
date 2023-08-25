package main

import controller "github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/Controller"

func main() {
	r := controller.GinRouter()
	r.Run()
}
