package feragstring

import "fmt"

type RouteInfo struct {
	FeragMessage
	routeName                 string
	limit                     int
	maxStack                  int
	standard                  int
	parameterN                int
	maxBundle                 int
	parameterSz               int
	topsheetMarker            int
	eaMarker                  int
	eaAddressDefinition       int
	topsheetTemplateDirectory int
	editionName               string
	productReferenceNumbers   []int
}

func (ri *RouteInfo) ParameterSz() string {
	return fmt.Sprintf("+35%04d", ri.parameterSz)
}

func (ri *RouteInfo) SetParameterSz(parameterSz int) {
	ri.parameterSz = parameterSz
}

func (ri *RouteInfo) ProductReferenceNumbers() string {
	var prreffmt string
	for _, pr := range ri.productReferenceNumbers {
		prreffmt += fmt.Sprintf("+41%02d", pr)
	}
	return prreffmt
}

func (ri *RouteInfo) AddProductReferenceNumber(productReferenceNumber int) {
	ri.productReferenceNumbers = append(ri.productReferenceNumbers, productReferenceNumber)
}

func (ri *RouteInfo) TopsheetTemplateDirectory() string {
	return fmt.Sprintf("+56%03d", ri.topsheetTemplateDirectory)
}

func (ri *RouteInfo) SetTopsheetTemplateDirectory(topsheetTemplateDirectory int) {
	ri.topsheetTemplateDirectory = topsheetTemplateDirectory
}

func (ri *RouteInfo) EaMarker() string {
	return fmt.Sprintf("+69%1d", ri.eaMarker)
}

func (ri *RouteInfo) SetEaMarker(eaMarker int) {
	ri.eaMarker = eaMarker
}

func (ri *RouteInfo) MaxBundle() string {
	return fmt.Sprintf("+34%04d", ri.maxBundle)
}

func (ri *RouteInfo) SetMaxBundle(maxBundle int) {
	ri.maxBundle = maxBundle
}

func (ri *RouteInfo) ParameterN() string {
	return fmt.Sprintf("+33%04d", ri.parameterN)
}

func (ri *RouteInfo) SetParameterN(parameterN int) {
	ri.parameterN = parameterN
}

func (ri *RouteInfo) Standard() string {
	return fmt.Sprintf("+32%04d", ri.standard)
}

func (ri *RouteInfo) SetStandard(standard int) {
	ri.standard = standard
}

func (ri *RouteInfo) MaxStack() string {
	return fmt.Sprintf("+31%04d", ri.maxStack)
}

func (ri *RouteInfo) SetMaxStack(maxStack int) {
	ri.maxStack = maxStack
}

func (ri *RouteInfo) Limit() string {
	return fmt.Sprintf("+30%04d", ri.limit)
}

func (ri *RouteInfo) SetLimit(limit int) {
	ri.limit = limit
}

func (ri *RouteInfo) EditionName() string {
	if ri.editionName == "" {
		return ""
	}
	return fmt.Sprintf("+20%-30s", ri.editionName)
}

func (ri *RouteInfo) SetEditionName(editionName string) {
	ri.editionName = editionName
}

func (ri *RouteInfo) EaAddressDefinition() string {
	if ri.eaAddressDefinition == 0 {
		return ""
	}
	return fmt.Sprintf("+91%06d", ri.eaAddressDefinition)
}

func (ri *RouteInfo) SetEaAddressDefinition(eaAddressDefinition int) {
	ri.eaAddressDefinition = eaAddressDefinition
}

func (ri *RouteInfo) TopsheetMarker() string {
	return fmt.Sprintf("+59%1d", ri.topsheetMarker)
}

func (ri *RouteInfo) SetTopsheetMarker(topsheetMarker int) {
	ri.topsheetMarker = topsheetMarker
}

func (ri *RouteInfo) SetRouteName(routeName string) {
	ri.routeName = routeName
}

func (ri *RouteInfo) RouteName() string {
	return fmt.Sprintf("+11%-13s", ri.routeName)
}

func NewRouteInfo() *RouteInfo {
	ri := RouteInfo{
		FeragMessage: FeragMessage{
			messageStart: "2402",
			messageEnd:   "!",
		},
	}
	return &ri
}

func (ri *RouteInfo) Payload() string {
	data := ri.RouteName()
	data += ri.Limit()
	data += ri.MaxStack()
	data += ri.Standard()
	data += ri.ParameterN()
	data += ri.MaxBundle()
	data += ri.ParameterSz()
	data += ri.TopsheetMarker()
	data += ri.EaMarker()
	data += ri.EaAddressDefinition()
	data += ri.TopsheetTemplateDirectory()
	data += ri.EditionName()
	data += ri.ProductReferenceNumbers()
	return data
}

func (ri *RouteInfo) Message() string {
	message := ri.FeragMessage.MessageTemplate()
	return message(&ri.FeragMessage, ri.Payload())
}
