package handler

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/database"
	"github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/helper"
	"github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/model"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var uss model.UserSessionState
var vegitables model.Vegitable
var menus model.Menu
var channelSercret string
var accessToken string
var bot *linebot.Client
var err error
var errstr string

func init() {
	log.Println("aaaaaa")
	uss.InitUserSessionState()
	vegitables.MarshallCsv()
	menus.MarshallCsv()
	helper.LoadEnv()
	accessToken = os.Getenv("access_token")
	channelSercret = os.Getenv("channel_secret")
	bot, err = linebot.New(channelSercret, accessToken)
	if err != nil {
		panic(err)
	}

}

func Webhook(ctx *gin.Context) {

	events, err := bot.ParseRequest(ctx.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			ctx.JSON(400, gin.H{"message": "Invalid signature"})
		} else {
			ctx.JSON(500, gin.H{"message": "Internal server error"})
			panic(err)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch message.Text {
				case "今日の特産野菜":
					uss.UserSessionClear(event.Source.UserID)
					rand.Seed(time.Now().UnixNano())
					i := rand.Intn(6)
					_, err = bot.ReplyMessage(event.ReplyToken,
						linebot.NewTextMessage(vegitables[i].Name+"\n\n"+vegitables[i].Content),
						linebot.NewImageMessage(vegitables[i].ImagePlace, vegitables[i].ImagePlace)).Do()
					if err != nil {
						panic(err)
					}
				case "スケジュール":
					uss.UserSessionClear(event.Source.UserID)
					MessageText := helper.ReadText("../static/texts/GinozaEvent.txt")
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(MessageText)).Do()
					if err != nil {
						panic(err)
					}

				case "スタンプラリー":
					uss.UserSessionClear(event.Source.UserID)
					MessageText := helper.ReadText("../static/texts/StampRallyHelp.txt")
					StampDone := [4]string{"未押印", "未押印", "未押印", "未押印"}
					ustamprally, err := database.UserStampRallyProgress(event.Source.UserID)
					log.Println(ustamprally)
					for _, us := range ustamprally {
						switch us.StampId {
						case 1:
							StampDone[0] = "〇"
						case 2:
							StampDone[1] = "〇"
						case 3:
							StampDone[2] = "〇"
						case 4:
							StampDone[3] = "〇"
						}
					}

					if err != nil {
						panic(err)
					}

					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf(MessageText, StampDone[0], StampDone[1], StampDone[2], StampDone[3]))).Do()
					if err != nil {
						panic(err)
					} else {
						uss.UserStampRallySessionValidToTrue(event.Source.UserID)
					}
				case "リセット":
					uss.UserSessionClear(event.Source.UserID)
					err = database.UserStampRallyReset(event.Source.UserID)
					if err != nil {
						panic(err)
					}
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("スタンプラリーの進捗をリセットしました！")).Do()
					if err != nil {
						panic(err)
					}

				case "飲食店メニュー":
					if uss.IsUserMenuSessionValid(event.Source.UserID) {
						ms := uss.Status[event.Source.UserID].Ms
						if ms.GetSessionState() == 1 {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("以下の中からメニューを見たいお店の番号を選択してください！\n"+"\n1.COFFEE&BREAK GINOZA FARM LAB"+"\n2.大也丸"+"\n3.食堂まんじろう"+"\n4.喰遊"+"\n5.美ら天"+"\n6.てんぷす食堂"+"\n ")).Do()
							if err != nil {
								panic(err)
							}
						}
					} else {
						uss.UserMenuSessionValidToTrue(event.Source.UserID)
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("以下の中からメニューを見たいお店の番号を選択してください！\n"+"\n1.COFFEE&BREAK GINOZA FARM LAB"+"\n2.大也丸"+"\n3.食堂まんじろう"+"\n4.喰遊"+"\n5.美ら天"+"\n6.てんぷす食堂")).Do()
						if err != nil {
							panic(err)
						}
					}
				case "1":
					if uss.IsUserMenuSessionValid(event.Source.UserID) {
						ms := uss.Status[event.Source.UserID].Ms
						if ms.GetSessionState() == 1 {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(menus[0].ImagePlace, menus[0].ImagePlace)).Do()
							if err != nil {
								panic(err)
							}
						}
					} else if uss.IsUserStampSessionValid(event.Source.UserID) {
						err = database.UserStampRallyUpdate(event.Source.UserID, 1)
						if err != nil {
							errstr = err.Error()
						}
						if strings.Contains(errstr, "UNIQUE") {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("BBQのスタンプはすでに押されています！")).Do()
							if err != nil {
								log.Println(err)
							}
						} else if err != nil {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("予期しないエラー!")).Do()
							if err != nil {
								log.Println(err)
							}
						} else {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("スタンプラリーのBBQのスタンプを押しました!")).Do()
							if err != nil {
								log.Println(err)

							}
						}

					}
				case "2":
					if uss.IsUserMenuSessionValid(event.Source.UserID) {
						log.Println("menu")
						ms := uss.Status[event.Source.UserID].Ms
						if ms.GetSessionState() == 1 {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(menus[1].ImagePlace, menus[1].ImagePlace)).Do()
							if err != nil {
								panic(err)
							}
						}
					} else if uss.IsUserStampSessionValid(event.Source.UserID) {
						err = database.UserStampRallyUpdate(event.Source.UserID, 2)
						if err != nil {
							errstr = err.Error()
						}
						if strings.Contains(errstr, "UNIQUE") {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("フォトフレーム作り体験のスタンプはすでに押されています！")).Do()
							if err != nil {
								log.Println(err)

							}
						} else if err != nil {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("予期しないエラー!")).Do()

							if err != nil {
								log.Println(err)
							}
						} else {

							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("スタンプラリーのフォトフレーム作り体験のスタンプを押しました!")).Do()
							if err != nil {
								log.Println(err)

							}
						}
					}
				case "3":
					if uss.IsUserMenuSessionValid(event.Source.UserID) {
						ms := uss.Status[event.Source.UserID].Ms
						if ms.GetSessionState() == 1 {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(menus[2].ImagePlace, menus[2].ImagePlace)).Do()
							if err != nil {
								panic(err)
							}
						}
					} else if uss.IsUserStampSessionValid(event.Source.UserID) {
						err = database.UserStampRallyUpdate(event.Source.UserID, 3)
						if err != nil {
							errstr = err.Error()
						}
						if strings.Contains(errstr, "UNIQUE") {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("島草履彫り体験のスタンプはすでに押されています！")).Do()
							if err != nil {
								log.Println(err)
							}
						} else if err != nil {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("予期しないエラー!")).Do()
							if err != nil {
								log.Println(err)
							}
						} else {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("島草履彫り体験のスタンプを押しました!")).Do()
							if err != nil {
								log.Println(err)

							}
						}
					}
				case "4":
					if uss.IsUserMenuSessionValid(event.Source.UserID) {
						ms := uss.Status[event.Source.UserID].Ms
						if ms.GetSessionState() == 1 {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(menus[3].ImagePlace, menus[3].ImagePlace)).Do()
							if err != nil {
								panic(err)
							}
						}
					} else if uss.IsUserStampSessionValid(event.Source.UserID) {
						err = database.UserStampRallyUpdate(event.Source.UserID, 4)
						if err != nil {
							errstr = err.Error()
						}
						if strings.Contains(errstr, "UNIQUE") {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("シーサー色付け体験のスタンプはすでに押されています！")).Do()
							if err != nil {
								log.Println(err)
							}
						} else if err != nil {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("予期しないエラー!")).Do()
							if err != nil {
								log.Println(err)
							}
						} else {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("シーサー色付け体験のスタンプを押しました!")).Do()
							if err != nil {
								log.Println(err)

							}
						}
					}
				case "5":
					if uss.IsUserMenuSessionValid(event.Source.UserID) {
						ms := uss.Status[event.Source.UserID].Ms
						if ms.GetSessionState() == 1 {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(menus[4].ImagePlace, menus[4].ImagePlace)).Do()
							if err != nil {
								panic(err)
							}
						}
					}
				case "6":
					if uss.IsUserMenuSessionValid(event.Source.UserID) {
						ms := uss.Status[event.Source.UserID].Ms
						if ms.GetSessionState() == 1 {
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(menus[5].ImagePlace, menus[5].ImagePlace)).Do()
							if err != nil {
								panic(err)
							}
						}
					}
				}
			}
		}
	}
}
