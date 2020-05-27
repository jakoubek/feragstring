package feragstring

import (
	"fmt"
	"time"
)

type TitleInfo struct {
	FeragMessage
	printObjectName string
	titleName string
	publicationDate time.Time
	countryCode string
	printObjectColor string
	additionalInfo string
}

func (ti *TitleInfo) SetPrintObjectName(printObjectName string) {
	ti.printObjectName = printObjectName
}

func (ti *TitleInfo) AdditionalInfo() string {
	return fmt.Sprintf("+08%-50s", ti.additionalInfo)
}

func (ti *TitleInfo) SetAdditionalInfo(additionalInfo string) {
	ti.additionalInfo = additionalInfo
}

func (ti *TitleInfo) PrintObjectColor() string {
	return fmt.Sprintf("+94%-8s", ti.printObjectColor)
}

func (ti *TitleInfo) SetPrintObjectColor(printObjectColor string) {
	ti.printObjectColor = printObjectColor
}

func (ti *TitleInfo) CountryCode() string {
	return fmt.Sprintf("+97%-2s", ti.countryCode)
}

func (ti *TitleInfo) SetCountryCode(countryCode string) {
	ti.countryCode = countryCode
}

func (ti *TitleInfo) SetPublicationDate(publicationDateString string) {
	parsedDate, err := time.Parse(dateInputFormatISO, publicationDateString)
	if err != nil {
		panic(err)
	}
	ti.publicationDate = parsedDate
}

func (ti *TitleInfo) PublicationDate() string {
	return fmt.Sprintf("+95%-6s", ti.publicationDate.Format(dateOutputFormat))
}


func (ti *TitleInfo) Message() string {
	message := ti.FeragMessage.MessageTemplate()
	return message(&ti.FeragMessage, ti.Payload())
}

func (ti *TitleInfo) Payload() string {
	data := ti.PrintObjectName()
	data += ti.TitleName()
	data += ti.PublicationDate()
	data += ti.CountryCode()
	data += ti.PrintObjectColor()
	data += ti.AdditionalInfo()
	return data
}

func (ti *TitleInfo) PrintObjectName() string {
	return fmt.Sprintf("+93%-12s", ti.printObjectName)
}

func (ti *TitleInfo) TitleName() string {
	return fmt.Sprintf("+40%-8s", ti.titleName)
}

func (ti *TitleInfo) SetTitleName(titleName string) {
	ti.titleName = titleName
}

func NewTitleInfo() *TitleInfo {
	t := TitleInfo{
		FeragMessage: FeragMessage{"2440", "!"},
	}
	return &t
}

