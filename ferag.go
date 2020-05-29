package feragstring

const (
	dateInputFormatISO = "2006-01-02"
	dateOutputFormat   = "060102"
	linebreak          = "\r\n"
)

type FeragString struct {
	TitleInfo             *TitleInfo
	TitleEnd              *TitleEnd
	ProductReferences     []*ProductReference
	ProductReferencesNr   int
	RouteListEntries      []*RouteListEntry
	RouteListEntryCount   int
	RouteInfoEntries      []*RouteInfo
	ProductionDropEntries []*ProductionDrop
	RouteEndEntries       []*RouteEnd
}

func NewFeragString() *FeragString {
	fs := FeragString{
		TitleInfo: NewTitleInfo(),
		TitleEnd:  NewTitleEnd(),
	}
	return &fs
}

func (fs *FeragString) SetTitleName(titleName string) {
	fs.TitleInfo.SetTitleName(titleName)
	fs.TitleEnd.SetTitleName(titleName)
}

func (fs *FeragString) PrintOut() string {
	// +2440 | first message is the title info
	info := fs.TitleInfo.Message()

	// +2450 | list of product references
	for _, pr := range fs.ProductReferences {
		info += pr.Message()
	}

	for _, rl := range fs.RouteListEntries {
		info += rl.Message()
	}

	for _, ri := range fs.RouteInfoEntries {
		info += ri.Message()
	}

	for _, pd := range fs.ProductionDropEntries {
		info += pd.Message()
	}

	for _, re := range fs.RouteEndEntries {
		info += re.Message()
	}

	// +2441 | last message is the corresponding title end
	info += fs.TitleEnd.Message()
	return info
}

func (fs *FeragString) AddProductReference(pr *ProductReference) error {
	fs.ProductReferencesNr++
	pr.SetProductReferenceNumber(fs.ProductReferencesNr)
	if pr.productReferenceNumber == 1 && pr.productUsageType == 0 {
		pr.SetProductUsageType(1)
	}
	fs.ProductReferences = append(fs.ProductReferences, pr)
	return nil
}

func (fs *FeragString) AddRouteListEntry(rl *RouteListEntry) error {
	fs.RouteListEntryCount++
	//if rl.routeCode == 0 {
	//	rl.SetRouteCode(fs.RouteListEntryCount)
	//}
	fs.RouteListEntries = append(fs.RouteListEntries, rl)
	return nil
}

func (fs *FeragString) NextRouteCode() int {
	return fs.RouteListEntryCount + 1
}

func (fs *FeragString) AddRouteInfo(ri *RouteInfo) error {
	fs.RouteInfoEntries = append(fs.RouteInfoEntries, ri)
	return nil
}

func (fs *FeragString) AddProductionDrop(pd *ProductionDrop) error {
	fs.ProductionDropEntries = append(fs.ProductionDropEntries, pd)
	return nil
}

func (fs *FeragString) AddRouteEnd(re *RouteEnd) error {
	fs.RouteEndEntries = append(fs.RouteEndEntries, re)
	return nil
}
