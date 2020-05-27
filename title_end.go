package feragstring

import "fmt"

type TitleEnd struct {
	FeragMessage
	titleName string
}

func (ti *TitleEnd) PrintObjectName() string {
	return fmt.Sprintf("+40%-8s", ti.titleName)
}

func (ti *TitleEnd) SetTitleName(titleName string) {
	ti.titleName = titleName
}

func NewTitleEnd() *TitleEnd {
	t := TitleEnd{
		FeragMessage: FeragMessage{
			messageStart: "2441",
			messageEnd:   "!",
		},
	}
	return &t
}

func (ti *TitleEnd) Message() string {
	message := ti.FeragMessage.MessageTemplate()
	payload := fmt.Sprintf("+40%-8s", ti.titleName)
	return message(&ti.FeragMessage, payload)
}