package commands

import (
	model "activity_tracker_bot/models"
	"activity_tracker_bot/services/whatsapp/utils"
	"fmt"
	"strings"
	"time"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"gorm.io/gorm"
)

// HandleDone handles /done command
func HandleList(db *gorm.DB, client *whatsmeow.Client, message *events.Message, args []string) {
	// check if this is a group or private message
	chatJID := message.Info.Chat
	group := &model.Group{}

	groupResult := db.Find(&group, "group_jid = ?", chatJID.String())

	if groupResult.RowsAffected == 0 {
		utils.SendReply(client, chatJID, "⚠️ Group is not registered!")
		return
	}

	//get user group
	var usersGroup []model.UserGroup
	db.Find(&usersGroup, "group_jid = ?", chatJID.String())

	if len(usersGroup) == 0 {
		utils.SendReply(client, chatJID, "⚠️ No user registered!")
		return
	}

	var lines []string
	for _, users := range usersGroup {
		var user model.User
		db.First(&user, "id = ?", users.UserID)
		lines = append(lines, fmt.Sprintf("👤 *%s*", user.Name))

		var userActivities []struct {
			ActivityName string
			UserName     string
			Status       string
		}

		uaResult := db.Table("user_activities").
			Select(`
				activities.name AS activity_name,
				periods.name AS period_name,
				routines.status as status
			`).
			Joins("JOIN activities ON user_activities.activity_id = activities.id").
			Joins("JOIN periods ON activities.period_id = periods.id").
			Joins("LEFT JOIN routines ON routines.user_id = user_activities.user_id AND routines.activity_id = activities.id").
			Where("user_activities.group_jid = ?", chatJID).
			Where("user_activities.user_id = ?", users.UserID).
			Find(&userActivities)
		if uaResult.Error != nil {
			utils.SendReply(client, chatJID, "⚠️ Failed to get user activities!")
			return
		}
		for _, ua := range userActivities {
			var status string
			switch ua.Status {
			case "done":
				status = "✅"
			case "not_done":
				status = "❌"
			default:
				status = "❌"
			}
			lines = append(lines, fmt.Sprintf("- %s %s", ua.ActivityName, status))

		}
	}

	today := time.Now()
	header := fmt.Sprintf(`📅 %s
Selamat pagi! 🌞
Hari baru adalah peluang baru untuk tumbuh dan lebih baik! 🌱
Jangan berkecil hati, perjalanan kebaikan dimulai dari langkah kecil yang terus diulang. 💫

📋 Checklist Ibadah Hari Ini:
`, utils.FormatTanggalIndo(today))

	footer := `📖 Renungan Hari Ini:
"Sesungguhnya Allah tidak akan mengubah keadaan suatu kaum sebelum mereka mengubah keadaan diri mereka sendiri."
(QS Ar-Ra'd: 11)

🔥 Masih ada waktu hari ini untuk mengejar dan memperbaiki. Yuk, semangat untuk lebih baik dari kemarin!`

	routinesList := strings.Join(lines, `
	`)

	text := fmt.Sprintf(`%s
	
	%s
	
	%s
	`, header, routinesList, footer)

	utils.SendReply(client, chatJID, text)
}
