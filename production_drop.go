package feragstring

import "fmt"

// ProductionDrop is the struct for one production drop
// underneath a route
type ProductionDrop struct {
	FeragMessage
	agentName         string
	numberOfCopies    int
	ControlCharacters ControlCharacterSet
	dontProduce       bool
	topsheetData      string
}

// TopsheetData returns the formatted topsheet data segment
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

// SetTopsheetData sets the topsheet data to a given string
func (pd *ProductionDrop) SetTopsheetData(topsheetData string) {
	pd.topsheetData = topsheetData
}

func (pd *ProductionDrop) SetDontProduce() {
	pd.dontProduce = true
}

// NumberOfCopies returns the formatted number of copies in the route
func (pd *ProductionDrop) NumberOfCopies() string {
	return fmt.Sprintf("+13%05d", pd.numberOfCopies)
}

// SetNumberOfCopies sets the number of copies in the production drop
func (pd *ProductionDrop) SetNumberOfCopies(numberOfCopies int) {
	pd.numberOfCopies = numberOfCopies
}

// AgentName returns the formatted agent name
func (pd *ProductionDrop) AgentName() string {
	return fmt.Sprintf("+12%-10s", pd.agentName)
}

// SetAgentName sets the agent name to a given string
func (pd *ProductionDrop) SetAgentName(agentName string) {
	pd.agentName = agentName
}

// NewProductionDrop instantiates a new production drop
// struct and returns a pointer to it.
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

// Payload returns the formatted FERAG string
// for embedding in the message
func (pd *ProductionDrop) Payload() string {
	data := pd.AgentName()
	data += pd.NumberOfCopies()
	data += pd.ControlCharacters.GetControlCharactersMessage()
	return data
}

// Message returns the formatted FERAG string
// for the production drop
func (pd *ProductionDrop) Message() string {
	message := pd.FeragMessage.MessageTemplate()
	return message(&pd.FeragMessage, pd.Payload())
}
