# Getting started

`structmap` is a Go library for converting struct and map to string.

Go offers a lot of ways of converting struct or map to string. This library is just another instant way of converting struct or map to string. It doesn't use a ready-to-use method from another 3rd party library, it uses raw methods instead.

## Installation

```
$ go get github.com/s-wijaya/structmap
```

## Import

```golang
import "github.com/s-wijaya/structmap"
```

## Usage

For this struct and its sample value:

```golang
type (
	FirstLayer struct {
		FirstString string
		FirstLayer  SecondLayer
	}
	SecondLayer struct {
		SecondString string
		SecondLayer  []ThirdLayer
	}
	ThirdLayer struct {
		ThirdString string
		ThirdLayer  FourthLayer
	}
	FourthLayer struct {
		FourthLayer int
	}
)

var sampleStruct = FirstLayer{
	FirstString: "first",
	FirstLayer: SecondLayer{
		SecondString: "second",
		SecondLayer: []ThirdLayer{
			{
				ThirdString: "third-one",
				ThirdLayer: FourthLayer{
					FourthLayer: 0,
				},
			},
			{
				ThirdString: "third-two",
				ThirdLayer: FourthLayer{
					FourthLayer: 1,
				},
			},
		},
	},
}
```

You can convert `sampleStruct` above to string using this simple code:
```golang
resultString := structmap.StructToString(sampleStruct)

// {"FirstString":"first","FirstLayer":{"SecondString":"second","SecondLayer":[{"ThirdString":"third-one","ThirdLayer":{"FourthLayer":0}},{"ThirdString":"third-two","ThirdLayer":{"FourthLayer":1}}]}}
```