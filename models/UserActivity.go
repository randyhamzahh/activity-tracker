package model

type UserActivity struct {
	ActivityID int `json:"activity_id" binding:"required"`
	UserID     int `json:"user_id" binding:"required"`
}
