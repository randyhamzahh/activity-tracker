package model

type UserActivity struct {
	ActivityID int    `json:"activity_id"`
	GroupJid   string `json:"group_jid"`
	GroupID    int    `json:"group_id"`
	UserID     int    `json:"user_id"`
	UserJid    string `json:"user_jid"`
}
