package main

import (
	"log"
	"os"
	"time"

	"github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/helper"
	"github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/model"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/robfig/cron/v3"
)

var GEvents model.Event
var channelSercret string
var accessToken string
var err error
var bot *linebot.Client

func init() {
	helper.LoadEnv()
	accessToken = os.Getenv("access_token")
	channelSercret = os.Getenv("channel_secret")
	bot, err = linebot.New(channelSercret, accessToken)
	if err != nil {
		panic(err)
	}

	GEvents.MarshallCsv()
}

func main() {
	/*c := cron.New()
	c.AddFunc("@every 1s", func() {
		day := time.Now().Weekday()
		log.Println(GEvents[day])
	})

	c.Start()*/

	c := cron.New()

	c.AddFunc("* * * * * ", func() {
		day := time.Now().Weekday()
		log.Println(int(day))
		log.Println(GEvents)
		if GEvents[int(day)].Content == "" {
			_, err = bot.BroadcastMessage(linebot.NewTextMessage("今日のイベントはありません")).Do()
			if err != nil {
				log.Println(err)
			}

		} else {
			log.Println("bbbbb")
			log.Println(GEvents[int(day)].Content)
			_, err = bot.BroadcastMessage(linebot.NewTextMessage("今日のイベントは「" + GEvents[int(day)].Content + "」です！")).Do()
			if err != nil {
				log.Println(err)
			}
		}
	})

	c.Start()
	time.Sleep(2 * time.Minute)

}
