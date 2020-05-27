package feragstring

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNewFeragString(t *testing.T) {
	fs := NewFeragString()
	fs.SetTitleName("MOP10629")

	fs.TitleInfo.SetPrintObjectName("MOP1")
	fs.TitleInfo.SetPublicationDate("2019-06-29")
	fs.TitleInfo.SetCountryCode("13")
	fs.TitleInfo.SetPrintObjectColor("03368448")

	// Produktreferenzen
	pr1 := NewProductReference()
	pr1.SetProductName("HP MOP1 KERN")
	pr1.SetCopiesAssigned(56433)
	pr1.SetSupervision(1)
	pr1.SetOverlap(5)
	mp := MissingParameter{
		missingRate: 1,
		missingSize: 1,
	}
	pr1.SetMissingParameter(mp)
	pr1.SetIssueReference("HPMOP1 K")
	fs.AddProductReference(pr1)
	// ------------------------------------------
	pr2 := NewProductReference()
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
	rl1 := NewRouteListEntry()
	rl1.SetRouteName("UT002223")
	rl1.SetCopiesInRoute(309)
	fs.AddRouteListEntry(rl1)
	rl2 := NewRouteListEntry()
	rl2.SetRouteName("UT888888")// Route Lists - Ende
	rl2.SetCopiesInRoute(80)
	fs.AddRouteListEntry(rl2)
	producedContent := fs.PrintOut()

	filename := "D:\\TEMP\\Feragstring\\ferag.txt"
	err := ioutil.WriteFile(filename, []byte(producedContent), 0644)
	if err != nil {
		panic(err)
	}

	testcontent := getTestFileContent("D:\\TEMP\\Feragstring\\test.txt")
	fmt.Println("=======================")
	fmt.Println("TEST:")
	fmt.Println(testcontent)
	fmt.Println("-----------------------")
	fmt.Println("PRODUCED:")
	fmt.Println(producedContent)
	fmt.Println("=======================")
	if testcontent != producedContent {
		t.Errorf("Files don't match!")
	}
}

func getTestFileContent(filename string) string {
	testcontent, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(testcontent)
}