package utils

import "github.com/google/uuid"

func GenUUID() string {
	return uuid.NewString()
}
