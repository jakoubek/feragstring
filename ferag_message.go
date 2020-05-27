package feragstring

import "fmt"

type FeragMessage struct {
	messageStart string
	messageEnd string
}

func (fm *FeragMessage) getMessageStart() string {
	return fmt.Sprintf("%%%s", fm.messageStart)
}

func (fm *FeragMessage) getMessageEnd() string {
	return fmt.Sprintf("%s", fm.messageEnd)
}

type MessageTemplateFunc func(*FeragMessage, string) string

func (fm *FeragMessage) MessageTemplate() MessageTemplateFunc {
	return func(fm *FeragMessage, s string) string {
		message := fm.getMessageStart()
		message += s
		message += fm.getMessageEnd() + linebreak
		return message
	}
}
