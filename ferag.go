package feragstring

const (
	dateInputFormatISO = "2006-01-02"
	dateOutputFormat   = "060102"
	linebreak          = "\r\n"
)

// FeragString is the root struct for creating
// a FERAG string. One FeragString instance returns
// the output for one FERAG string.
type FeragString struct {
	TitleInfo             *TitleInfo
	TitleEnd              *TitleEnd
	ProductReferences     []*ProductReference
	ProductReferencesNr   int
	Routes                []*Route
	RouteCount            int
	RouteListEntries      []*RouteListEntry
	RouteListEntryCount   int
	RouteInfoEntries      []*RouteInfo
	ProductionDropEntries []*ProductionDrop
	RouteEndEntries       []*RouteEnd
}

// NewFeragString instantiates a new FeragString
// struct and returns a pointer to it.
func NewFeragString() *FeragString {
	fs := FeragString{
		TitleInfo: NewTitleInfo(),
		TitleEnd:  NewTitleEnd(),
	}
	return &fs
}

// SetTitleName sets the title name field
func (fs *FeragString) SetTitleName(titleName string) {
	fs.TitleInfo.SetTitleName(titleName)
	fs.TitleEnd.SetTitleName(titleName)
}

// PrintOut returns the final FERAG string of the FeragString instance
func (fs *FeragString) PrintOut() string {
	// +2440 | first message is the title info
	info := fs.TitleInfo.Message()

	// +2450 | list of product references
	for _, pr := range fs.ProductReferences {
		info += pr.Message()
	}

	// create route list entries for every route
	for _, rt := range fs.Routes {
		info += rt.GetRouteListEntry().Message()
	}

	// create route info for every route
	// including embedded production drops
	for _, rt := range fs.Routes {
		info += rt.GetRouteMessage()
	}

	//for _, rl := range fs.RouteListEntries {
	//	info += rl.Message()
	//}
	//
	//for _, ri := range fs.RouteInfoEntries {
	//	info += ri.Message()
	//}
	//
	//for _, pd := range fs.ProductionDropEntries {
	//	info += pd.Message()
	//}
	//
	//for _, re := range fs.RouteEndEntries {
	//	info += re.Message()
	//}

	// +2441 | last message is the corresponding title end
	info += fs.TitleEnd.Message()
	return info
}

// AddProductReference adds a Product Reference instance to the list
// of product references of the FeragString
func (fs *FeragString) AddProductReference(pr *ProductReference) error {
	if pr.productReferenceNumber == 0 {
		fs.ProductReferencesNr++
		pr.SetProductReferenceNumber(fs.ProductReferencesNr)
	}
	if pr.productReferenceNumber == 1 && pr.productUsageType == 0 {
		pr.SetProductUsageType(1)
	}
	fs.ProductReferences = append(fs.ProductReferences, pr)
	return nil
}

// AddRoute adds a Route to the list of routes of the FeragString
func (fs *FeragString) AddRoute(r *Route) error {
	fs.RouteCount++
	fs.Routes = append(fs.Routes, r)
	return nil
}

// NextRouteCode returns the next numeric route code.
// Use this method to get an automatically incremented
// counter value for the routes of a FeragString.
func (fs *FeragString) NextRouteCode() int {
	return fs.RouteCount + 1
	//return fs.RouteListEntryCount + 1
}

// Deprecated: AddRouteListEntry adds a route list entry to a FeragString.
// Use AddRoute instead.
func (fs *FeragString) AddRouteListEntry(rl *RouteListEntry) error {
	fs.RouteListEntryCount++
	//if rl.routeCode == 0 {
	//	rl.SetRouteCode(fs.RouteListEntryCount)
	//}
	fs.RouteListEntries = append(fs.RouteListEntries, rl)
	return nil
}

// Deprecated: AddRouteInfo adds a route info to a FeragString.
// Use AddRoute instead.
func (fs *FeragString) AddRouteInfo(ri *RouteInfo) error {
	fs.RouteInfoEntries = append(fs.RouteInfoEntries, ri)
	return nil
}

// Deprecated: AddProductionDrop adds a ProductionDrop to a FeragString.
// Use AddProductionDrop on the route struct instead.
func (fs *FeragString) AddProductionDrop(pd *ProductionDrop) error {
	fs.ProductionDropEntries = append(fs.ProductionDropEntries, pd)
	return nil
}

// Deprecated: AddRouteEnd adds a Route End to a FeragString.
// Not necessary anymore. Routes are automatically ended now.
func (fs *FeragString) AddRouteEnd(re *RouteEnd) error {
	fs.RouteEndEntries = append(fs.RouteEndEntries, re)
	return nil
}
