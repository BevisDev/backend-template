package utils

import (
	"encoding/json"
	"reflect"
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

func CheckTypeAndConvert(value interface{}) interface{} {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr ||
		v.Kind() == reflect.Struct ||
		v.Kind() == reflect.Map ||
		v.Kind() == reflect.Slice ||
		v.Kind() == reflect.Array {
		return ToJSON(value)
	}
	return value
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
