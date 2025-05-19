package utils

import (
	"context"
	"fmt"
	"log"
	"net/mail"
	"strings"
	"time"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	waTypes "go.mau.fi/whatsmeow/types"
)

// SendReply sends a simple text message to a user or group
func SendReply(client *whatsmeow.Client, chatJID waTypes.JID, text string) {

	_, err := client.SendMessage(context.Background(), chatJID, &waProto.Message{
		Conversation: protoString(text),
	})

	if err != nil {
		log.Printf("❌ Failed to send response: %v", err)
	}
}

func IsGroupChat(jid waTypes.JID) bool {
	return jid.Server == "g.us"
}

// helper
func protoString(s string) *string {
	return &s
}

func GetPhoneNumberFromJID(jid types.JID) string {
	rawJID := jid.String()
	parts := strings.Split(rawJID, "@")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func InArray(target string, list string) bool {
	for _, item := range strings.Split(list, ",") {
		if item == target {
			return true
		}
	}
	return false
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

var hariIndo = map[time.Weekday]string{
	time.Sunday:    "Minggu",
	time.Monday:    "Senin",
	time.Tuesday:   "Selasa",
	time.Wednesday: "Rabu",
	time.Thursday:  "Kamis",
	time.Friday:    "Jumat",
	time.Saturday:  "Sabtu",
}

var bulanIndo = map[time.Month]string{
	time.January:   "Januari",
	time.February:  "Februari",
	time.March:     "Maret",
	time.April:     "April",
	time.May:       "Mei",
	time.June:      "Juni",
	time.July:      "Juli",
	time.August:    "Agustus",
	time.September: "September",
	time.October:   "Oktober",
	time.November:  "November",
	time.December:  "Desember",
}

// FormatTanggalIndo returns a string like "🗓️ Kamis, 27 Maret 2025"
func FormatTanggalIndo(t time.Time) string {
	hari := hariIndo[t.Weekday()]
	bulan := bulanIndo[t.Month()]
	return fmt.Sprintf("%s, %02d %s %d", hari, t.Day(), bulan, t.Year())
}
