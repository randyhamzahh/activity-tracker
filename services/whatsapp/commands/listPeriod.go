package commands

import (
	model "activity_tracker_bot/models"
	"activity_tracker_bot/services/whatsapp/utils"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"gorm.io/gorm"
)

// HandleDone handles /done command
func HandleListPeriod(db *gorm.DB, client *whatsmeow.Client, message *events.Message, args []string) {

	// check if this is a group or private message
	// check if this is a group  and if group is already registered
	chatJID := message.Info.Chat
	group := &model.Group{}
	var period []model.Period

	groupResult := db.Find(&group, "group_jid = ?", chatJID.String())

	if groupResult.RowsAffected == 0 {
		utils.SendReply(client, chatJID, "⚠️ Group is not registered!")
		return
	}

	result := db.Find(&period)

	if result.Error != nil {
		utils.SendReply(client, chatJID, "⚠️ Failed to get periods!")
		return
	}

	// check if group is already registered
	var text string
	for _, p := range period {
		text += "- " + p.Name + "\n"
	}

	utils.SendReply(client, chatJID, text)
}
