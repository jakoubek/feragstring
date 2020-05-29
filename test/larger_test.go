package test

import (
	"fmt"
	"io/ioutil"
	"jakoubek.net/feragstring"
	"testing"
)

func TestNewFeragString(t *testing.T) {
	fs := feragstring.NewFeragString()
	fs.SetTitleName("MOP10629")

	fs.TitleInfo.SetPrintObjectName("MOP1")
	fs.TitleInfo.SetPublicationDate("2019-06-29")
	fs.TitleInfo.SetCountryCode("13")
	fs.TitleInfo.SetPrintObjectColor("03368448")
	fs.TitleInfo.ShowEmptyAdditionalInfo()

	// Produktreferenzen
	pr1 := feragstring.NewProductReference()
	pr1.SetProductName("HP MOP1 KERN")
	pr1.SetCopiesAssigned(56433)
	pr1.SetSupervision(1)
	pr1.SetOverlap(5)
	mp := feragstring.NewMissingParameter(1, 1)
	pr1.SetMissingParameter(mp)
	pr1.SetIssueReference("HPMOP1 K")
	fs.AddProductReference(pr1)
	// ------------------------------------------
	pr2 := feragstring.NewProductReference()
	pr2.SetProductName("Beilage Rossmann ET 29.06.19")
	pr2.SetProductUsageType(3)
	pr2.SetCopiesAssigned(36000)
	pr2.SetProductWeight(180)
	pr2.SetSupervision(1)
	pr2.SetOverlap(1)
	pr2.SetIssueReference("19004378")
	fs.AddProductReference(pr2)
	// Produktreferenzen - Ende

	// Route Lists
	rl1 := feragstring.NewRouteListEntry()
	rl1.SetRouteName("UT002223")
	rl1.SetRouteCode(fs.NextRouteCode())
	rl1.SetRampNumber(0)
	rl1.SetCopiesInRoute(309)
	fs.AddRouteListEntry(rl1)

	rl2 := feragstring.NewRouteListEntry()
	rl2.SetRouteName("UT888888")
	rl2.SetRouteCode(fs.NextRouteCode())
	rl2.SetRampNumber(0)
	rl2.SetCopiesInRoute(80)
	fs.AddRouteListEntry(rl2)
	// Route Lists - Ende

	producedContent := fs.PrintOut()

	filename := "D:\\TEMP\\Feragstring\\larger_test.txt"
	err := ioutil.WriteFile(filename, []byte(producedContent), 0644)
	if err != nil {
		panic(err)
	}

	testcontent := getTestFileContent("D:\\TEMP\\Feragstring\\larger_want.txt")
	fmt.Println("=======================")
	fmt.Println("WANT:")
	fmt.Println(testcontent)
	fmt.Println("-----------------------")
	fmt.Println("PRODUCED:")
	fmt.Println(producedContent)
	fmt.Println("=======================")
	if testcontent != producedContent {
		t.Errorf("Files don't match!")
	}
}