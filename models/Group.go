package model

type Group struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserId int    `json:"user_id"`
}
