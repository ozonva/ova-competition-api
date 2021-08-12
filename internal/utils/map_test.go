package utils

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

type ValuesToKeysCase struct {
	originalMap   map[int]string
	reversedMap   map[string]int
	expectedError bool
}

func TestValuesToKeys(t *testing.T) {
	testCases := []ValuesToKeysCase{
		{
			nil, nil, true,
		},
		{
			map[int]string{}, map[string]int{}, false,
		},
		{
			map[int]string{1: "a", 2: "b", 3: "c"}, map[string]int{"a": 1, "b": 2, "c": 3}, false,
		},
		{
			map[int]string{1: "a", 2: "a", 3: "c"}, nil, true,
		},
	}

	for _, testCase := range testCases {
		reversedMap, err := ValuesToKeys(testCase.originalMap)
		if err != nil {
			if !testCase.expectedError {
				t.Fatal("inverting error", err)
			}
		}

		if !cmp.Equal(reversedMap, testCase.reversedMap) {
			t.Fatal("maps do not match", reversedMap, testCase.reversedMap)
		}
	}
}
