package model

import "time"

type Routine struct {
	ID         int       `json:"id"`
	Date       time.Time `json:"date"`
	UserID     uint      `json:"user_id"`
	ActivityID int       `json:"activity_id"`
	PeriodID   int       `json:"period_id"`
	Status     string    `json:"status"`
}
