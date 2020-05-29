package feragstring

import "fmt"

type RouteListEntry struct {
	FeragMessage
	routeName string
	routeCode int
	rampNumber int
	copiesInRoute int
}

func (r *RouteListEntry) CopiesInRoute() string {
	if r.copiesInRoute == 0 {
		return ""
	}
	return fmt.Sprintf("+23%06d", r.copiesInRoute)
}

func (r *RouteListEntry) SetCopiesInRoute(copiesInRoute int) {
	r.copiesInRoute = copiesInRoute
}

func (r *RouteListEntry) RampNumber() string {
	if r.rampNumber == 0 {
		return ""
	}
	return fmt.Sprintf("+25%02d", r.rampNumber)
}

func (r *RouteListEntry) SetRampNumber(rampNumber int) {
	r.rampNumber = rampNumber
}

func (r *RouteListEntry) RouteCode() string {
	if r.routeCode == 0 {
		return ""
	}
	return fmt.Sprintf("+79%05d", r.routeCode)
}

func (r *RouteListEntry) SetRouteCode(routeCode int) {
	r.routeCode = routeCode
}

func (r *RouteListEntry) RouteName() string {
	return fmt.Sprintf("+11%-13s", r.routeName)
}

func (r *RouteListEntry) SetRouteName(routeName string) {
	r.routeName = routeName
}

func NewRouteListEntry() *RouteListEntry {
	rl := RouteListEntry{
		FeragMessage: FeragMessage{
			messageStart: "2401",
			messageEnd:   "!",
		},
	}
	return &rl
}

func (rl *RouteListEntry) Payload() string {
	data := rl.RouteName()
	data += rl.RouteCode()
	data += rl.RampNumber()
	data += rl.CopiesInRoute()
	return data
}

func (rl *RouteListEntry) Message() string {
	message := rl.FeragMessage.MessageTemplate()
	return message(&rl.FeragMessage, rl.Payload())
}
