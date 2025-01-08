package utils

import (
	"context"
	"errors"
	"reflect"
	"strings"
)

func IsNilOrEmpty(inp interface{}) bool {
	if inp == nil {
		return true
	}
	v := reflect.ValueOf(inp)
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.String:
		return strings.TrimSpace(v.String()) == ""
	case reflect.Array, reflect.Slice:
		return v.Len() == 0
	case reflect.Map:
		return v.Len() == 0
	case reflect.Chan:
		return v.Len() == 0
	case reflect.Ptr:
		return v.IsNil()
	default:
		return false
	}
}

func IsPtr(inp interface{}) bool {
	return reflect.ValueOf(inp).Kind() == reflect.Ptr
}

func IsTimedOut(err error) bool {
	return errors.Is(err, context.DeadlineExceeded)
}
