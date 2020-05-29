package feragstring

import (
	"fmt"
	"strings"
)

// FeragMessage is the struct for the "abstract" message.
// It is meant to be embedded into the real message structs.
type FeragMessage struct {
	messageStart string
	messageEnd   string
}

func (fm *FeragMessage) getMessageStart() string {
	return fmt.Sprintf("%%%s", fm.messageStart)
}

func (fm *FeragMessage) getMessageEnd() string {
	return fmt.Sprintf("%s", fm.messageEnd)
}

// MessageTemplateFunc is a func type for the generating the
// FERAG-formatted message.
type MessageTemplateFunc func(*FeragMessage, string) string

// MessageTemplate returns a MessageTemplateFunc - to be called
// from within the "real" message structs.
func (fm *FeragMessage) MessageTemplate() MessageTemplateFunc {
	return func(fm *FeragMessage, s string) string {
		message := fm.getMessageStart()
		message += s
		message += fm.getMessageEnd()
		return strings.TrimSpace(message) + linebreak
	}
}
