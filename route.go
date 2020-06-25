package feragstring

type Route struct {
	routeName                 string
	routeCode                 int
	rampNumber                int
	copiesInRoute             int
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
	productReferenceNumber    int
	ProductionDrops           []*ProductionDrop
}

func (r *Route) GetRouteListEntry() *RouteListEntry {
	rle := NewRouteListEntry()
	rle.SetRouteName(r.routeName)
	rle.SetRouteCode(r.routeCode)
	rle.SetRampNumber(r.rampNumber)
	rle.SetCopiesInRoute(r.copiesInRoute)
	return rle
}

func (r *Route) GetRouteMessage() string {
	ri := NewRouteInfo()
	ri.SetRouteName(r.routeName)
	info := ri.Message()

	for _, pd := range r.ProductionDrops {
		info += pd.Message()
		info += pd.TopsheetData()
	}

	re := NewRouteEnd()
	re.SetRouteName(r.routeName)
	info += re.Message()

	return info
}

func (r *Route) AddProductionDrop(pd *ProductionDrop) error {
	r.ProductionDrops = append(r.ProductionDrops, pd)
	return nil
}

func (r *Route) SetRouteName(routeName string) {
	r.routeName = routeName
}

func (r *Route) SetRouteCode(routeCode int) {
	r.routeCode = routeCode
}

func (r *Route) SetRampNumber(rampNumber int) {
	r.rampNumber = rampNumber
}

func (r *Route) SetCopiesInRoute(copiesInRoute int) {
	r.copiesInRoute = copiesInRoute
}

func (ri *Route) SetLimit(limit int) {
	ri.limit = limit
}

func (ri *Route) SetMaxStack(maxStack int) {
	ri.maxStack = maxStack
}

func (ri *Route) SetStandard(standard int) {
	ri.standard = standard
}

func (ri *Route) SetParameterN(parameterN int) {
	ri.parameterN = parameterN
}

func (ri *Route) SetMaxBundle(maxBundle int) {
	ri.maxBundle = maxBundle
}

func (ri *Route) SetTopsheetMarker(topsheetMarker int) {
	ri.topsheetMarker = topsheetMarker
}

func (ri *Route) SetEaMarker(eaMarker int) {
	ri.eaMarker = eaMarker
}

func (ri *Route) SetTopsheetTemplateDirectory(topsheetTemplateDirectory int) {
	ri.topsheetTemplateDirectory = topsheetTemplateDirectory
}

func (ri *Route) SetProductReferenceNumber(productReferenceNumber int) {
	ri.productReferenceNumber = productReferenceNumber
}

func NewRoute() *Route {
	r := Route{
		rampNumber: -1,
	}
	return &r
}