package commands

import (
	model "activity_tracker_bot/models"
	"activity_tracker_bot/services/whatsapp/utils"
	"log"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"gorm.io/gorm"
)

// HandleDone handles /done command
func HandleRegGroup(db *gorm.DB, client *whatsmeow.Client, message *events.Message, args []string) {
	chatJID := message.Info.Chat
	// check if this is a group or private message
	if !utils.IsGroupChat(chatJID) {
		utils.SendReply(client, chatJID, "⚠️ Please register in group chat!")
		return
	}

	group := &model.Group{}

	db.Find(&group, "group_jid = ?", chatJID.String())

	if group.ID != 0 {
		utils.SendReply(client, chatJID, "⚠️ Group already registered!")
		return
	}

	groupInfo, err := client.GetGroupInfo(chatJID)

	if err != nil {
		log.Printf("❌ Failed to get group info: %v", err)
		return
	}

	group = &model.Group{
		Name:     groupInfo.Name,
		GroupJid: chatJID.String(),
	}

	groupCreated := db.Create(&group)

	if groupCreated.Error != nil {
		utils.SendReply(client, chatJID, "⚠️ Failed to register group!")
		return
	}
	utils.SendReply(client, chatJID, "✅ Group registered!")
}
