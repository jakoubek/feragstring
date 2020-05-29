package feragstring

import "fmt"

type ProductReference struct {
	FeragMessage
	productReferenceNumber int
	productName            string
	productUsageType       int
	sheetLayers            int
	copiesAssigned         int
	productWeight          int
	supervision            int
	overlap                int
	feedingMode            int
	scatter                int
	missingParameter       MissingParameter
	issueReference         string
}

type MissingParameter struct {
	missingRate int
	missingSize int
}

func NewMissingParameter(missingRate, missingSize int) MissingParameter {
	return MissingParameter{
		missingRate: missingRate,
		missingSize: missingSize,
	}
}

func (pr *ProductReference) IssueReference() string {
	return fmt.Sprintf("+99195%-8s", pr.issueReference)
}

func (pr *ProductReference) SetIssueReference(issueReference string) {
	pr.issueReference = issueReference
}

func (pr *ProductReference) MissingParameter() string {
	return fmt.Sprintf("+99105%04d%08d", pr.missingParameter.missingRate, pr.missingParameter.missingSize)
}

func (pr *ProductReference) SetMissingParameter(missingParameter MissingParameter) {
	pr.missingParameter = missingParameter
}

func (pr *ProductReference) Scatter() string {
	return fmt.Sprintf("+99102%06d", pr.scatter)
}

func (pr *ProductReference) SetScatter(scatter int) {
	pr.scatter = scatter
}

func (pr *ProductReference) FeedingMode() string {
	return fmt.Sprintf("+99101%02d", pr.feedingMode)
}

func (pr *ProductReference) SetFeedingMode(feedingMode int) {
	pr.feedingMode = feedingMode
}

func (pr *ProductReference) Overlap() string {
	return fmt.Sprintf("+45%02d", pr.overlap)
}

func (pr *ProductReference) SetOverlap(overlap int) {
	pr.overlap = overlap
}

func (pr *ProductReference) Supervision() string {
	return fmt.Sprintf("+44%02d", pr.supervision)
}

func (pr *ProductReference) SetSupervision(supervision int) {
	pr.supervision = supervision
}

func (pr *ProductReference) ProductWeight() string {
	return fmt.Sprintf("+36%05d", pr.productWeight)
}

func (pr *ProductReference) SetProductWeight(productWeight int) {
	pr.productWeight = productWeight
}

func (pr *ProductReference) CopiesAssigned() string {
	return fmt.Sprintf("+02%08d", pr.copiesAssigned)
}

func (pr *ProductReference) SetCopiesAssigned(copiesAssigned int) {
	pr.copiesAssigned = copiesAssigned
}

func (pr *ProductReference) SheetLayers() string {
	return fmt.Sprintf("+35%04d", pr.sheetLayers)
}

func (pr *ProductReference) SetSheetLayers(sheetLayers int) {
	pr.sheetLayers = sheetLayers
}

func (pr *ProductReference) ProductUsageType() string {
	return fmt.Sprintf("+55%02d", pr.productUsageType)
}

func (pr *ProductReference) SetProductUsageType(productUsageType int) {
	pr.productUsageType = productUsageType
}

func (pr *ProductReference) ProductName() string {
	return fmt.Sprintf("+42%-30s", pr.productName)
}

func (pr *ProductReference) SetProductName(productName string) {
	pr.productName = productName
}

func (pr *ProductReference) ProductReferenceNumber() string {
	return fmt.Sprintf("+41%02d", pr.productReferenceNumber)
}

func (pr *ProductReference) SetProductReferenceNumber(productReferenceNumber int) {
	pr.productReferenceNumber = productReferenceNumber
}

func NewProductReference() *ProductReference {
	pr := ProductReference{
		FeragMessage:     FeragMessage{"2450", "!"},
		missingParameter: MissingParameter{
			missingRate: 0,
			missingSize: 1,
		},
	}
	return &pr
}

func (pr *ProductReference) Payload() string {
	data := pr.ProductReferenceNumber()
	data += pr.ProductName()
	data += pr.ProductUsageType()
	data += pr.SheetLayers()
	data += pr.CopiesAssigned()
	data += pr.ProductWeight()
	data += pr.Supervision()
	data += pr.Overlap()
	data += pr.FeedingMode()
	data += pr.MissingParameter()
	data += pr.IssueReference()
	return data
}

func (pr *ProductReference) Message() string {
	message := pr.FeragMessage.MessageTemplate()
	return message(&pr.FeragMessage, pr.Payload())
}