package commands

import (
	model "activity_tracker_bot/models"
	"activity_tracker_bot/services/whatsapp/utils"
	"fmt"
	"time"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"gorm.io/gorm"
)

// HandleDone handles /done command
func HandleDone(db *gorm.DB, client *whatsmeow.Client, message *events.Message, args []string) {

	// check if this is a group or private message
	chatJID := message.Info.Chat
	group := &model.Group{}
	senderJID := message.Info.Sender
	// senderName := message.Info.PushName
	user := &model.User{}
	activity := args[0]

	groupResult := db.Find(&group, "group_jid = ?", chatJID.String())

	// check if group is already registered
	if groupResult.RowsAffected == 0 {
		utils.SendReply(client, chatJID, "⚠️ Group is not registered!")
		return
	}

	// check if user is already registered
	db.Find(&user, "user_jid = ?", senderJID.String())
	if user.ID == 0 {
		utils.SendReply(client, chatJID, "⚠️ User is not registered!")
		return
	}

	// ceck if reg activity format
	if len(args) < 1 {
		utils.SendReply(client, message.Info.Chat, "⚠️ Please use the correct format:\n/done;<activity>")
		return
	}

	// check if activity is already registered
	var activityModel model.Activity
	db.Find(&activityModel, "name = ?", activity)
	var userActivities []model.UserActivity
	db.Find(&userActivities, "user_id = ? AND activity_id = ?", user.ID, activityModel.ID)
	if activityModel.ID == 0 && userActivities[0].ActivityID == 0 {
		utils.SendReply(client, chatJID, "⚠️ Activity is not registered!")
		return
	}

	// check if routine is marked as done
	var routine model.Routine
	db.Find(&routine, "date = ? AND user_id = ? AND activity_id = ?", time.Now().Format("2006-01-02"), user.ID, activityModel.ID)
	if routine.ID != 0 {
		if routine.Status == "done" {
			utils.SendReply(client, chatJID, fmt.Sprintf("⚠️ %s is already marked as done!", activityModel.Name))
			return
		}
		routine.Status = "done"
		db.Save(&routine)
	}
	// save routine
	routineModel := model.Routine{
		Date:       time.Now(), // or any other format you prefer
		UserID:     user.ID,
		ActivityID: activityModel.ID,
		PeriodID:   activityModel.PeriodID,
		Status:     "done",
	}

	db.Create(&routineModel)
	// reply to group
	utils.SendReply(client, chatJID, fmt.Sprintf("✅ %s is marked as done!", activityModel.Name))
}
