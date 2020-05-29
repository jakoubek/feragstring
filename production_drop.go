package feragstring

import "fmt"

type ProductionDrop struct {
	FeragMessage
	agentName string
	numberOfCopies int
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
	}
	return &pd
}

func (pd *ProductionDrop) Payload() string {
	data := pd.AgentName()
	data += pd.NumberOfCopies()
	return data
}

func (pd *ProductionDrop) Message() string {
	message := pd.FeragMessage.MessageTemplate()
	return message(&pd.FeragMessage, pd.Payload())
}
