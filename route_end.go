package feragstring

import "fmt"

// RouteEnd is the struct that automatically ends a route definition
type RouteEnd struct {
	FeragMessage
	routeName string
}

// RouteName returns the formatted route name segment
func (re *RouteEnd) RouteName() string {
	return fmt.Sprintf("+11%-13s", re.routeName)
}

// SetRouteName sets the route name field
func (re *RouteEnd) SetRouteName(routeName string) {
	re.routeName = routeName
}

// NewRouteEnd  instantiates a new Route End
// struct and returns a pointer to it.
func NewRouteEnd() *RouteEnd {
	re := RouteEnd{
		FeragMessage: FeragMessage{
			messageStart: "2406",
			messageEnd:   "!",
		},
	}
	return &re
}

// Payload returns the formatted FERAG string
// for embedding in the message
func (re *RouteEnd) Payload() string {
	data := re.RouteName()
	return data
}

// Message returns the formatted FERAG string
// for the complete route end message
func (re *RouteEnd) Message() string {
	message := re.FeragMessage.MessageTemplate()
	return message(&re.FeragMessage, re.Payload())
}
