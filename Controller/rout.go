package controller

import (
	"github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/Controller/handler"
	"github.com/gin-gonic/gin"
)

func GinRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//LineBotのWebhookのエンドポイント
	r.POST("/webhook", handler.Webhook)

	return r

}
