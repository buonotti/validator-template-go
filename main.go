package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func fail(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

func validate(testCase ValidationItem) {
	// TODO: validate test case and use fail() to fail the test case.
}

func main() {
	strIn := ""
	_, err := fmt.Scanln(&strIn)
	if err != nil {
		fail("Could not read input: " + err.Error())
	}
	var item ValidationItem
	err = json.Unmarshal([]byte(strIn), &item)
	if err != nil {
		fail("Could not deserialize input: " + err.Error())
	}
	validate(item)
}
