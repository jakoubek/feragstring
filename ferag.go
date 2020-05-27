package feragstring

const (
	dateInputFormatISO = "2006-01-02"
	dateOutputFormat = "060102"
	linebreak = "\r\n"
)

type FeragString struct {
	TitleInfo *TitleInfo
	TitleEnd *TitleEnd
	ProductReferences []*ProductReference
	ProductReferencesNr int
	RouteListEntries []*RouteListEntry
	RouteListEntryCount int
}

func NewFeragString() *FeragString {
	fs := FeragString{
		TitleInfo: NewTitleInfo(),
		TitleEnd: NewTitleEnd(),
	}
	return &fs
}

func (fs *FeragString) SetTitleName(titleName string) {
	fs.TitleInfo.SetTitleName(titleName)
	fs.TitleEnd.SetTitleName(titleName)
}

func (fs *FeragString) PrintOut() string {
	info := fs.TitleInfo.Message()

	for _, pr := range fs.ProductReferences {
		info += pr.Message()
	}

	for _, rl := range fs.RouteListEntries {
		info += rl.Message()
	}

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
	if rl.routeCode == 0 {
		rl.SetRouteCode(fs.RouteListEntryCount)
	}
	fs.RouteListEntries = append(fs.RouteListEntries, rl)
	return nil
}