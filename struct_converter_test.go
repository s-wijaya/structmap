package structmap

import (
	"testing"
)

type (
	FirstLayer struct {
		FirstComp string
	}
)

var sampleStruct = FirstLayer{
	FirstComp: "first",
}

var expectedResultString = "{\"FirstComp\":\"first\"}"

func TestStructToString(test *testing.T) {
	test.Logf("String result : %s", StructToString(sampleStruct))

	if string(StructToString(sampleStruct)) != string(expectedResultString) {
		test.Errorf("Expected string result : %s", expectedResultString)
	}
}
