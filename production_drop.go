package feragstring

import "fmt"

// ProductionDrop is the struct for one production drop
// underneath a route
type ProductionDrop struct {
	FeragMessage
	agentName               string
	numberOfCopies          int
	ControlCharacters       ControlCharacterSet
	limit                   int
	maxStack                int
	standard                int
	parameterN              int
	maxBundle               int
	parameterSz             int
	dontProduce             bool
	topsheetData            string
	productReferenceNumbers []int
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

func (pd *ProductionDrop) MaxBundle() string {
	return fmt.Sprintf("+34%04d", pd.maxBundle)
}

func (pd *ProductionDrop) SetMaxBundle(maxBundle int) {
	pd.maxBundle = maxBundle
}

func (pd *ProductionDrop) ParameterN() string {
	return fmt.Sprintf("+33%04d", pd.parameterN)
}

func (pd *ProductionDrop) SetParameterN(parameterN int) {
	pd.parameterN = parameterN
}

func (pd *ProductionDrop) Standard() string {
	return fmt.Sprintf("+32%04d", pd.standard)
}

func (pd *ProductionDrop) SetStandard(standard int) {
	pd.standard = standard
}

func (pd *ProductionDrop) MaxStack() string {
	return fmt.Sprintf("+31%04d", pd.maxStack)
}

func (pd *ProductionDrop) SetMaxStack(maxStack int) {
	pd.maxStack = maxStack
}

func (pd *ProductionDrop) Limit() string {
	return fmt.Sprintf("+30%04d", pd.limit)
}

func (pd *ProductionDrop) SetLimit(limit int) {
	pd.limit = limit
}

func (pd *ProductionDrop) ParameterSz() string {
	return fmt.Sprintf("+35%04d", pd.parameterSz)
}

func (pd *ProductionDrop) SetParameterSz(parameterSz int) {
	pd.parameterSz = parameterSz
}

// SetTopsheetData sets the topsheet data to a given string
func (pd *ProductionDrop) SetTopsheetData(topsheetData string) {
	pd.topsheetData = topsheetData
}

// ProductReferenceNumbers returns the string of TSL-formatted ProductReferenceNumbers
func (pd *ProductionDrop) ProductReferenceNumbers() string {
	var prreffmt string
	for _, pr := range pd.productReferenceNumbers {
		prreffmt += fmt.Sprintf("+99141%03d", pr)
	}
	return prreffmt
}

// AddProductReferenceNumber adds a numeric ProductReferenceNumber to the production drop
func (pd *ProductionDrop) AddProductReferenceNumber(productReferenceNumber int) {
	pd.productReferenceNumbers = append(pd.productReferenceNumbers, productReferenceNumber)
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
	data += pd.Limit()
	data += pd.MaxStack()
	data += pd.Standard()
	data += pd.ParameterN()
	data += pd.MaxBundle()
	data += pd.ParameterSz()
	data += pd.ProductReferenceNumbers()
	return data
}

// Message returns the formatted FERAG string
// for the production drop
func (pd *ProductionDrop) Message() string {
	message := pd.FeragMessage.MessageTemplate()
	return message(&pd.FeragMessage, pd.Payload())
}
