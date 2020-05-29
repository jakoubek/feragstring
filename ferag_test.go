package feragstring

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestShortestFeragString(t *testing.T) {
	fs := NewFeragString()
	fs.SetTitleName("DEMO2009")

	rl := NewRouteListEntry()
	rl.SetRouteName("E1_ROUTE_100")
	fs.AddRouteListEntry(rl)

	ri := NewRouteInfo()
	ri.SetRouteName("E1_ROUTE_100")
	ri.SetEditionName("E1")
	fs.AddRouteInfo(ri)

	pd := NewProductionDrop()
	pd.SetAgentName("R100RE001")
	pd.SetNumberOfCopies(123)
	fs.AddProductionDrop(pd)

	re := NewRouteEnd()
	re.SetRouteName("E1_ROUTE_100")
	fs.AddRouteEnd(re)

	producedContent := fs.PrintOut()

	want := getTestFileContent("D:\\TEMP\\Feragstring\\minimal_want.txt")

	fmt.Println("--WANT-----------------")
	fmt.Println(want)
	fmt.Println("-----------------------")
	fmt.Println("--PRODUCED-------------")
	fmt.Println(producedContent)
	fmt.Println("-----------------------")
	fmt.Println("=======================")

	filename := "D:\\TEMP\\Feragstring\\minimal_test.txt"
	err := ioutil.WriteFile(filename, []byte(producedContent), 0644)
	if err != nil {
		panic(err)
	}


	fmt.Printf("COMPARE: %d", strings.Compare(producedContent, want))
	if strings.Compare(producedContent, want) != 0 {
	//if strings.TrimSpace(producedContent) != strings.TrimSpace(want) {
		t.Errorf("Produced result does not equal to minimal example")
	}

}

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

func getTestFileContent(filename string) string {
	testcontent, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(testcontent)
}