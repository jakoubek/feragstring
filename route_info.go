package feragstring

import "fmt"

type RouteInfo struct {
	FeragMessage
	routeName           string
	topsheetMarker      int
	eaAddressDefinition int
	editionName         string
}

func (ri *RouteInfo) EditionName() string {
	return fmt.Sprintf("+20%-30s", ri.editionName)
}

func (ri *RouteInfo) SetEditionName(editionName string) {
	ri.editionName = editionName
}

func (ri *RouteInfo) EaAddressDefinition() string {
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
	data += ri.TopsheetMarker()
	data += ri.EaAddressDefinition()
	data += ri.EditionName()
	return data
}

func (ri *RouteInfo) Message() string {
	message := ri.FeragMessage.MessageTemplate()
	return message(&ri.FeragMessage, ri.Payload())
}
