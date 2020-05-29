package feragstring

import "fmt"

type RouteEnd struct {
	FeragMessage
	routeName string
}

func (re *RouteEnd) RouteName() string {
	return fmt.Sprintf("+11%-13s", re.routeName)
}

func (re *RouteEnd) SetRouteName(routeName string) {
	re.routeName = routeName
}

func NewRouteEnd() *RouteEnd {
	re := RouteEnd{
		FeragMessage:        FeragMessage{
			messageStart: "2406",
			messageEnd:   "!",
		},
	}
	return &re
}

func (re *RouteEnd) Payload() string {
	data := re.RouteName()
	return data
}

func (re *RouteEnd) Message() string {
	message := re.FeragMessage.MessageTemplate()
	return message(&re.FeragMessage, re.Payload())
}
