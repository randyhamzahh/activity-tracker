package initializers

import (
	"activity_tracker_bot/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

var DB *gorm.DB
var BotSession *sqlstore.Container

func buildPgDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DatabaseConfig.DBHost,
		config.DatabaseConfig.DBUser,
		config.DatabaseConfig.DBPassword,
		config.DatabaseConfig.DBName,
		config.DatabaseConfig.DBPort,
		config.DatabaseConfig.SSLMode,
	)
}

func ConnectToPgSql() *gorm.DB {
	var err error

	DB, err = gorm.Open(postgres.Open(buildPgDSN()), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return DB
}

func ConnectToWhatsAppSession() *sqlstore.Container {

	BotSession, err := sqlstore.New("postgres", buildPgDSN(), nil)
	if err != nil {
		log.Fatalf("Failed to connect WhatsApp session store: %v", err)
	}

	return BotSession
}
