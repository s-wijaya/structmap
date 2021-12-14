package structmap

import (
	"bytes"
	"encoding/json"
	"strings"
)

func StructToString(data interface{}) string {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	stringResult := reqBodyBytes.String()
	trimmedString := strings.ReplaceAll(stringResult, " ", "")

	return trimmedString[:len(trimmedString)-1]
}
