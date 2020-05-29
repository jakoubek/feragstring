package test

import (
	"io/ioutil"
	"jakoubek.net/feragstring"
	"strings"
	"testing"
)

func TestShortestFeragString(t *testing.T) {
	fs := feragstring.NewFeragString()
	fs.SetTitleName("DEMO2009")

	rl := feragstring.NewRouteListEntry()
	rl.SetRouteName("E1_ROUTE_100")
	fs.AddRouteListEntry(rl)

	ri := feragstring.NewRouteInfo()
	ri.SetRouteName("E1_ROUTE_100")
	ri.SetEditionName("E1")
	fs.AddRouteInfo(ri)

	pd := feragstring.NewProductionDrop()
	pd.SetAgentName("R100RE001")
	pd.SetNumberOfCopies(123)
	fs.AddProductionDrop(pd)

	re := feragstring.NewRouteEnd()
	re.SetRouteName("E1_ROUTE_100")
	fs.AddRouteEnd(re)

	producedContent := fs.PrintOut()

	want := getTestFileContent("D:\\TEMP\\Feragstring\\minimal_want.txt")

	//fmt.Println("--WANT-----------------")
	//fmt.Println(want)
	//fmt.Println("-----------------------")
	//fmt.Println("--PRODUCED-------------")
	//fmt.Println(producedContent)
	//fmt.Println("-----------------------")
	//fmt.Println("=======================")

	filename := "D:\\TEMP\\Feragstring\\minimal_test.txt"
	err := ioutil.WriteFile(filename, []byte(producedContent), 0644)
	if err != nil {
		panic(err)
	}

	if strings.Compare(producedContent, want) != 0 {
		//if strings.TrimSpace(producedContent) != strings.TrimSpace(want) {
		t.Errorf("Produced result does not equal to minimal example")
	}
}
