package feragstring

import "fmt"

// TitleEnd is the struct for the FERAG title info message (%2441)
type TitleEnd struct {
	FeragMessage
	titleName string
}

// TitleName returns the title name segment (+40) FERAG-formatted
func (ti *TitleEnd) TitleName() string {
	return fmt.Sprintf("+40%-8s", ti.titleName)
}

// SetTitleName sets the title name segment (+40)
func (ti *TitleEnd) SetTitleName(titleName string) {
	ti.titleName = titleName
}

// NewTitleEnd returns a new TitleEnd struct
func NewTitleEnd() *TitleEnd {
	t := TitleEnd{
		FeragMessage: FeragMessage{
			messageStart: "2441",
			messageEnd:   "!",
		},
	}
	return &t
}

// Message returns the complete FERAG-formatted message for title end
func (ti *TitleEnd) Message() string {
	message := ti.FeragMessage.MessageTemplate()
	payload := ti.TitleName()
	return message(&ti.FeragMessage, payload)
}
