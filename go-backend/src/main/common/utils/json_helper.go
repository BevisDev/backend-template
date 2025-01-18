package utils

import (
	"encoding/json"
)

func ToJSON(v any) []byte {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return []byte("{}")
	}
	return jsonBytes
}

func ToJSONStr(v any) string {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(jsonBytes)
}

func ToStruct(jsonBytes []byte, result any) error {
	err := json.Unmarshal(jsonBytes, result)
	if err != nil {
		return err
	}
	return nil
}

func FromJSONBytes(jsonBytes []byte) string {
	return string(jsonBytes)
}

func FromJSONStr(jsonStr string, result any) error {
	err := json.Unmarshal([]byte(jsonStr), result)
	if err != nil {
		return err
	}
	return nil
}
