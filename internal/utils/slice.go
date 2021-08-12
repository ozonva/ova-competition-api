package utils

import (
	"errors"
	"fmt"
)

// IntSliceToBatches converts given slice to multiple slices of given size
func IntSliceToBatches(slice []int, chunkSize int) ([][]int, error) {
	if slice == nil {
		return nil, errors.New("received nil slice")
	}

	if chunkSize <= 0 {
		return nil, errors.New(fmt.Sprintf("chunk size should be positive, got %d", chunkSize))
	}

	batches := make([][]int, 0)
	for i := 0; i <= len(slice); i += chunkSize {
		sliceEnd := i + chunkSize
		if sliceEnd > len(slice) {
			sliceEnd = len(slice)
		}

		if sliceEnd-i > 0 {
			batches = append(batches, slice[i:sliceEnd])
		}
	}

	return batches, nil
}

func FilterWords(slice []string, words []string) []string {
	wordsSet := make(map[string]struct{})
	for _, word := range words {
		wordsSet[word] = struct{}{}
	}

	filtered := make([]string, 0)
	for _, value := range slice {
		if _, ok := wordsSet[value]; !ok {
			filtered = append(filtered, value)
		}
	}

	return filtered
}
