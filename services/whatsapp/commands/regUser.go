package commands

import (
	model "activity_tracker_bot/models"
	"activity_tracker_bot/services/whatsapp/utils"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"gorm.io/gorm"
)

// HandleDone handles /done command
func HandleRegUser(db *gorm.DB, client *whatsmeow.Client, message *events.Message, args []string) {

	chatJID := message.Info.Chat
	senderJID := message.Info.Sender
	senderName := message.Info.PushName
	group := &model.Group{}
	user := &model.User{}
	email := args[0]
	password := args[1]

	groupResult := db.Find(&group, "group_jid = ?", chatJID.String())

	if groupResult.RowsAffected == 0 {
		utils.SendReply(client, chatJID, "⚠️ Group is not registered!")
		return
	}

	if len(args) < 2 {
		utils.SendReply(client, message.Info.Chat, "⚠️ Please use the correct format:\n/reg-user;<email>;<password>")
		return
	}

	if !utils.IsValidEmail(args[0]) {
		utils.SendReply(client, chatJID, "⚠️ Please use the correct email!")
		return

	}

	db.Find(&user, "user_jid = ?", senderJID.String())

	if user.ID != 0 {
		utils.SendReply(client, chatJID, "⚠️ User already registered!")
		return
	}

	user = &model.User{
		Name:     senderName,
		Email:    email,
		Password: password,
		Phone:    utils.GetPhoneNumberFromJID(senderJID),
		UserJid:  senderJID.String(),
	}

	userReslut := db.Create(&user)

	uaResult := db.Table("user_groups").Create(map[string]interface{}{
		"user_id":   user.ID,
		"user_jid":  senderJID.String(),
		"group_jid": chatJID.String(),
		"group_id":  group.ID,
	})

	if userReslut.Error != nil && uaResult.Error != nil {
		utils.SendReply(client, chatJID, "⚠️ Failed to register user!")
		return
	}

	utils.SendReply(client, chatJID, "✅ User registered successfully!")
}
