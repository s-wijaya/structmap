package structmap

import (
	"bytes"
	"encoding/json"
)

func StructToString(data interface{}) string {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	return reqBodyBytes.String()
}
