package model

type Group struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	GroupJid string `json:"group_jid"`
}
