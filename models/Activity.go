package model

type Activity struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PeriodID int    `json:"period_id"`
}
