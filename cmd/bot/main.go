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
}
