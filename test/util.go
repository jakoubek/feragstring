package test

import "io/ioutil"

func getTestFileContent(filename string) string {
	testcontent, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(testcontent)
}
