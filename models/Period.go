package model

type Period struct {
	ID           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	DaysOfWeek   string `json:"days_of_week" binding:"required"`
	DaysOfMonth  string `json:"days_of_month"`
	MonthsOfYear string `json:"months_of_year"`
}
