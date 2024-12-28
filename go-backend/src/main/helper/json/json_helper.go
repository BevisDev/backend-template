package json

import (
	"encoding/json"
)

func ToJSONStr(v any) string {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(jsonBytes)
}

func ToJSON(v any) []byte {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return []byte("{}")
	}
	return jsonBytes
}

func ToStruct(jsonBytes []byte, result any) error {
	err := json.Unmarshal(jsonBytes, result)
	if err != nil {
		return err
	}
	return nil
}

func FromJSONStr(jsonStr string, result any) error {
	err := json.Unmarshal([]byte(jsonStr), result)
	if err != nil {
		return err
	}
	return nil
}
