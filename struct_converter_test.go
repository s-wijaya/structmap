package structmap

import "testing"

type (
	FirstLayer struct {
		FirstComp string
		NextLayer SecondLayer
	}
	SecondLayer struct {
		SecondComp string
		NextLayer  ThirdLayer
	}
	ThirdLayer struct {
		ThirdComp string
	}
)

var sampleStruct = FirstLayer{
	FirstComp: "first",
	NextLayer: SecondLayer{
		SecondComp: "second",
		NextLayer: ThirdLayer{
			ThirdComp: "third",
		},
	},
}

var expectedResultString = "{\"FirstComp\":\"first\",\"NextLayer\":{\"SecondComp\":\"second\",\"NextLayer\":{\"ThirdComp\":\"third\"}}}"

func TestStructToString(test *testing.T) {
	test.Logf("String result : %s", StructToString(sampleStruct))

	if StructToString(sampleStruct) != expectedResultString {
		test.Errorf("Expected string result %s", expectedResultString)
	}
}
