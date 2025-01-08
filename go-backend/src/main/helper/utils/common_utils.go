package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math"
)

func GenUUID() string {
	return uuid.NewString()
}

func Batches[T any](source []T, length int) ([][]T, error) {
	if length < 0 {
		return nil, fmt.Errorf("invalid length = %d", length)
	}

	var result [][]T
	size := len(source)
	if size <= 0 {
		return result, nil
	}

	fullChunks := int(math.Ceil(float64(size) / float64(length)))

	for n := 0; n < fullChunks; n++ {
		start := n * length
		end := start + length
		if end > size {
			end = size
		}
		result = append(result, source[start:end])
	}

	return result, nil
}
