package json

import (
	"encoding/json"
)

func ToJSON(v interface{}) []byte {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return []byte("{}")
	}
	return jsonBytes
}

func ToStruct(jsonStr string, result interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), result)
	if err != nil {
		return err
	}
	return nil
}
