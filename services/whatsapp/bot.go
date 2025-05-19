package whatsapp

import (
	"activity_tracker_bot/initializers"
	"context"
	"fmt"
	"os"
	"os/signal"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

type WhatsAppBot struct {
	Client *whatsmeow.Client
}

func NewBot() (*WhatsAppBot, error) {
	sessionStore := initializers.ConnectToWhatsAppSession()

	deviceStore, err := sessionStore.GetFirstDevice()
	if err != nil {
		return nil, fmt.Errorf("no device found: %w", err)
	}

	client := whatsmeow.NewClient(deviceStore, nil)

	return &WhatsAppBot{Client: client}, nil
}

func (b *WhatsAppBot) Start() error {

	b.Client.AddEventHandler(func(evt interface{}) {
		fmt.Println("Event trigered.....")
		switch v := evt.(type) {
		case *events.Message:
			fmt.Printf("Event Message ........... %s", v)
			fmt.Println("..........")
			HandleMessage(b.Client, v)
		}
	})

	if b.Client.Store.ID == nil {
		qrChan, _ := b.Client.GetQRChannel(context.Background())
		go func() {
			for evt := range qrChan {
				if evt.Event == "code" {
					fmt.Println("QR Code:", evt.Code)
					qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)

				} else {
					fmt.Println("QR event:", evt.Event)
					qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)

				}
			}
		}()
	}

	err := b.Client.Connect()
	if err != nil {
		return err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Shutting down...")
	b.Client.Disconnect()
	return nil
}
