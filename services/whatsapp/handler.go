package whatsapp

import (
	"activity_tracker_bot/initializers"
	"activity_tracker_bot/services/whatsapp/commands"
	"activity_tracker_bot/services/whatsapp/utils"

	"fmt"
	"strings"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func HandleMessage(client *whatsmeow.Client, message *events.Message) {

	fmt.Println("Handle Message Trigered....")
	// Ignore non-text messages
	if message.Message.GetConversation() == "" {
		println("Has no conversation.........")
		return
	}

	msgText := message.Message.GetConversation()

	// Only handle messages that start with "/"
	if !strings.HasPrefix(msgText, "/") {
		println("Has no prefix..............")
		return
	}

	// Extract the command (e.g. /done drink water)
	command := strings.TrimPrefix(msgText, "/")
	parts := strings.Split(command, ";")
	if len(parts) == 0 {
		println("parts is zero")
		return
	}

	cmd := parts[0]
	args := parts[1:]
	println("Command: " + cmd)
	println("Args: " + strings.Join(args, ";"))
	println("Arg length: " + fmt.Sprint(len(args)))

	// args1 := parts[1]
	// args2 := parts[2]
	db := initializers.DB // use encapsulated DB

	switch cmd {
	case "reg-group":
		commands.HandleRegGroup(db, client, message, args)
	case "reg-user":
		commands.HandleRegUser(db, client, message, args)
	case "reg-activity":
		commands.HandleRegActivity(db, client, message, args)
	case "list-period":
		commands.HandleListPeriod(db, client, message, args)
	case "list":
		commands.HandleList(db, client, message, args)
	case "done":
		commands.HandleDone(db, client, message, args)
	default:
		utils.SendReply(client, message.Info.Chat, "⚠️ Unknown command: /"+cmd)
		return
	}
}
