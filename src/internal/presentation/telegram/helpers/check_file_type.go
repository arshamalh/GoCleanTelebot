package helpers

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"strings"
)

// CheckFileType to check is file type valid?
func CheckFileType(m *tb.Message) bool {
	fileName := strings.ToLower(m.Document.FileName)
	return strings.HasSuffix(fileName, ".xlsx") || strings.HasSuffix(fileName, ".csv")
}
