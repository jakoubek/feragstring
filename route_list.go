package feragstring

import "fmt"

// RouteListEntry is the struct for one route entry in the
// list of routes
type RouteListEntry struct {
	FeragMessage
	routeName     string
	routeCode     int
	rampNumber    int
	copiesInRoute int
}

// CopiesInRoute returns the formatted number of copies in the route
func (r *RouteListEntry) CopiesInRoute() string {
	if r.copiesInRoute == 0 {
		return ""
	}
	return fmt.Sprintf("+23%06d", r.copiesInRoute)
}

// SetCopiesInRoute set the number of copies in the route
func (r *RouteListEntry) SetCopiesInRoute(copiesInRoute int) {
	r.copiesInRoute = copiesInRoute
}

// RampNumber returns the formatted ramp number. Segment is
// returned only if the ramp number is set to a value.
func (r *RouteListEntry) RampNumber() string {
	if r.rampNumber == -1 {
		return ""
	}
	return fmt.Sprintf("+25%02d", r.rampNumber)
}

// SetRampNumber sets the ramp number field to an integer
// value. A value of 0 is allowed.
func (r *RouteListEntry) SetRampNumber(rampNumber int) {
	r.rampNumber = rampNumber
}

// RouteCode returns the formatted route code.
func (r *RouteListEntry) RouteCode() string {
	if r.routeCode == 0 {
		return ""
	}
	return fmt.Sprintf("+79%05d", r.routeCode)
}

// SetRouteCode sets the route code field
func (r *RouteListEntry) SetRouteCode(routeCode int) {
	r.routeCode = routeCode
}

// RouteName returns the formatted route name.
func (r *RouteListEntry) RouteName() string {
	return fmt.Sprintf("+11%-13s", r.routeName)
}

// SetRouteName sets the route name field
func (r *RouteListEntry) SetRouteName(routeName string) {
	r.routeName = routeName
}

// NewRouteListEntry instantiates a new Route List Entry
// struct and returns a pointer to it.
func NewRouteListEntry() *RouteListEntry {
	rl := RouteListEntry{
		FeragMessage: FeragMessage{
			messageStart: "2401",
			messageEnd:   "!",
		},
		rampNumber: -1,
	}
	return &rl
}

// Payload returns the formatted FERAG string
// for embedding in the message
func (r *RouteListEntry) Payload() string {
	data := r.RouteName()
	data += r.RouteCode()
	data += r.RampNumber()
	data += r.CopiesInRoute()
	return data
}

// Message returns the formatted FERAG string
// for the complete route list entry
func (r *RouteListEntry) Message() string {
	message := r.FeragMessage.MessageTemplate()
	return message(&r.FeragMessage, r.Payload())
}
