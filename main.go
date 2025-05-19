package main

import (
	"activity_tracker_bot/config"
	"activity_tracker_bot/initializers"
	"activity_tracker_bot/services/whatsapp"
	"fmt"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectToPgSql()
	initializers.ConnectToWhatsAppSession()
}

func main() {

	bot, err := whatsapp.NewBot()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting WhatsApp bot...")
	if err := bot.Start(); err != nil {
		panic(err)
	}

	// router := gin.New()
	// router.Use(gin.Logger())
	// routes.RegisterAPIRoutes(router)

	// router.GET("/ting", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"success": "tong"})
	// })

	// router.Run(":" + config.AppConfig.AppPort)
}
