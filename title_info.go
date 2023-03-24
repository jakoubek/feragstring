package feragstring

import (
	"fmt"
	"time"
)

// TitleInfo is the struct for the FERAG title info message (%2440)
type TitleInfo struct {
	FeragMessage
	printObjectName         string
	titleName               string
	publicationDate         time.Time
	issueReference          string
	countryCode             string
	printObjectColor        string
	additionalInfo          string
	showEmptyAdditionalInfo bool
}

// ShowEmptyAdditionalInfo sets the flag showEmptyAdditionalInfo to true.
// The segment additional info is then shown even it is empty (+08).
func (ti *TitleInfo) ShowEmptyAdditionalInfo() {
	ti.showEmptyAdditionalInfo = true
}

// SetPrintObjectName sets the printObjectName segment (+93)
func (ti *TitleInfo) SetPrintObjectName(printObjectName string) {
	ti.printObjectName = printObjectName
}

// AdditionalInfo returns the additional info segment (+08) FERAG-formatted.
// It is not returned if empty UNLESS the flag showEmptyAdditionalInfo is
// set to true.
func (ti *TitleInfo) AdditionalInfo() string {
	if ti.additionalInfo == "" && ti.showEmptyAdditionalInfo == false {
		return ""
	}
	return fmt.Sprintf("+08%-50s", ti.additionalInfo)
}

// SetAdditionalInfo sets the additional info segment (+08)
func (ti *TitleInfo) SetAdditionalInfo(additionalInfo string) {
	ti.additionalInfo = additionalInfo
}

// PrintObjectColor returns the print object color segment (+94) FERAG-formatted
func (ti *TitleInfo) PrintObjectColor() string {
	if ti.printObjectColor == "" {
		return ""
	}
	return fmt.Sprintf("+94%-8s", ti.printObjectColor)
}

// SetPrintObjectColor sets the print object color segment (+94)
func (ti *TitleInfo) SetPrintObjectColor(printObjectColor string) {
	ti.printObjectColor = printObjectColor
}

// CountryCode returns the country code segment (+97) FERAG-formatted
func (ti *TitleInfo) CountryCode() string {
	if ti.countryCode == "" {
		return ""
	}
	return fmt.Sprintf("+97%-2s", ti.countryCode)
}

// SetCountryCode set the country code segment (+97)
func (ti *TitleInfo) SetCountryCode(countryCode string) {
	ti.countryCode = countryCode
}

// SetPublicationDate sets the publication date segment (+95)
func (ti *TitleInfo) SetPublicationDate(publicationDateString string) {
	parsedDate, err := time.Parse(dateInputFormatISO, publicationDateString)
	if err != nil {
		panic(err)
	}
	ti.publicationDate = parsedDate
}

// PublicationDate returns the publication date segment (+95) FERAG-formatted
func (ti *TitleInfo) PublicationDate() string {
	if ti.publicationDate.IsZero() {
		return ""
	}
	return fmt.Sprintf("+95%-6s", ti.publicationDate.Format(dateOutputFormat))
}

// Message returns the complete FERAG-formatted message for title info
func (ti *TitleInfo) Message() string {
	message := ti.FeragMessage.MessageTemplate()
	return message(&ti.FeragMessage, ti.Payload())
}

// Payload returns the inner content for title info to be included in Message()
func (ti *TitleInfo) Payload() string {
	data := ti.PrintObjectName()
	data += ti.TitleName()
	data += ti.PublicationDate()
	data += ti.IssueReference()
	data += ti.CountryCode()
	data += ti.PrintObjectColor()
	data += ti.AdditionalInfo()
	return data
}

func (ti *TitleInfo) SetIssueReference(issueReference string) {
	ti.issueReference = issueReference
}

func (ti *TitleInfo) IssueReference() string {
	if ti.issueReference == "" {
		return ""
	}
	return fmt.Sprintf("+99195%8s", ti.issueReference)
}

// PrintObjectName returns the print object name segment (+93) FERAG-formatted
func (ti *TitleInfo) PrintObjectName() string {
	if ti.printObjectName == "" {
		return ""
	}
	return fmt.Sprintf("+93%-12s", ti.printObjectName)
}

// TitleName returns the title name segment (+40) FERAG-formatted
func (ti *TitleInfo) TitleName() string {
	return fmt.Sprintf("+40%-8.8s", ti.titleName)
}

// SetTitleName sets the title name segment (+40)
func (ti *TitleInfo) SetTitleName(titleName string) {
	ti.titleName = titleName
}

// NewTitleInfo returns a new TitleInfo struct
func NewTitleInfo() *TitleInfo {
	t := TitleInfo{
		FeragMessage:            FeragMessage{"2440", "!"},
		showEmptyAdditionalInfo: false,
	}
	return &t
}
