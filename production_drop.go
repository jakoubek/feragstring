package feragstring

import "fmt"

type ProductionDrop struct {
	FeragMessage
	agentName         string
	numberOfCopies    int
	ControlCharacters ControlCharacterSet
	dontProduce       bool
	topsheetData      string
}

func (pd *ProductionDrop) TopsheetData() string {
	if pd.topsheetData == "" {
		return ""
	}
	tsd := pd.topsheetData
	if len(tsd) > 5996 {
		tsd = tsd[:5996]
	}

	tsdSegment := fmt.Sprintf("+58%s", tsd)

	fm := FeragMessage{
		messageStart: "2414",
		messageEnd:   "!",
	}
	message := fm.MessageTemplate()
	return message(&fm, tsdSegment)
}

func (pd *ProductionDrop) SetTopsheetData(topsheetData string) {
	pd.topsheetData = topsheetData
}

func (pd *ProductionDrop) SetDontProduce() {
	pd.dontProduce = true
}

func (pd *ProductionDrop) NumberOfCopies() string {
	return fmt.Sprintf("+13%05d", pd.numberOfCopies)
}

func (pd *ProductionDrop) SetNumberOfCopies(numberOfCopies int) {
	pd.numberOfCopies = numberOfCopies
}

func (pd *ProductionDrop) AgentName() string {
	return fmt.Sprintf("+12%-10s", pd.agentName)
}

func (pd *ProductionDrop) SetAgentName(agentName string) {
	pd.agentName = agentName
}

func NewProductionDrop() *ProductionDrop {
	pd := ProductionDrop{
		FeragMessage: FeragMessage{
			messageStart: "2403",
			messageEnd:   "!",
		},
		ControlCharacters: ControlCharacterSet{},
	}
	return &pd
}

func (pd *ProductionDrop) Payload() string {
	data := pd.AgentName()
	data += pd.NumberOfCopies()
	data += pd.ControlCharacters.GetControlCharactersMessage()
	return data
}

func (pd *ProductionDrop) Message() string {
	message := pd.FeragMessage.MessageTemplate()
	return message(&pd.FeragMessage, pd.Payload())
}
