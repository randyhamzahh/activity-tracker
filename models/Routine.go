package model

type Routine struct {
	ID         int    `json:"id"`
	Date       string `json:"date"`
	UserID     int    `json:"user_id"`
	ActivityID int    `json:"activity_id"`
	PeriodID   int    `json:"period_id"`
	Status     string `json:"status"`
}
