package commands

import (
	model "activity_tracker_bot/models"
	"activity_tracker_bot/services/whatsapp/utils"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"gorm.io/gorm"
)

// HandleDone handles /done command
func HandleRegActivity(db *gorm.DB, client *whatsmeow.Client, message *events.Message, args []string) {
	// check if this is a group  and if group is already registered
	chatJID := message.Info.Chat
	group := &model.Group{}
	activityName := args[0]
	periodName := args[1]
	activity := &model.Activity{}

	groupResult := db.Find(&group, "group_jid = ?", chatJID.String())

	if groupResult.RowsAffected == 0 {
		utils.SendReply(client, chatJID, "⚠️ Group is not registered!")
		return
	}
	// check if reg activity format
	if len(args) < 2 {
		utils.SendReply(client, message.Info.Chat, "⚠️ Please use the correct format:\n/reg-activity;<activity-name>;<period>")
		return
	}

	var period model.Period
	db.Find(&period, "name = ?", periodName)
	if period.ID == 0 {
		utils.SendReply(client, chatJID, "⚠️ Period is not exist!")
		return
	}

	var existingActivity model.Activity
	err := db.
		Joins("JOIN user_activities ON user_activities.activity_id = activities.id").
		Where("user_activities.group_jid = ?", chatJID.String()).
		Where("activities.name = ?", activityName).
		First(&existingActivity).Error

	if err == nil && existingActivity.ID != 0 {
		utils.SendReply(client, chatJID, "⚠️ Activity already exists!")
		return
	}

	activity = &model.Activity{
		Name:     activityName,
		PeriodID: period.ID,
	}

	result := db.Create(&activity)

	var userGroup []model.UserGroup

	ugResult := db.Table("user_groups").
		Where("group_jid = ?", chatJID).
		Find(&userGroup)

	for _, u := range userGroup {
		ua := &model.UserActivity{
			UserID:     u.UserID,
			UserJid:    u.UserJid,
			GroupID:    u.GroupID,
			GroupJid:   u.GroupJid,
			ActivityID: activity.ID,
		}
		db.Create(&ua)
	}

	if result.Error != nil && ugResult.Error != nil {
		utils.SendReply(client, chatJID, "⚠️ Failed to create activity!")
		return
	}

	utils.SendReply(client, chatJID, "✅ Activity created successfully!")
}
