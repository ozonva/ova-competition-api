package utils

import (
	"errors"
	"fmt"
)

// ValuesToKeys converts input map (int -> string) to map (string -> int)
func ValuesToKeys(inputMap map[int]string) (map[string]int, error) {
	if inputMap == nil {
		return nil, errors.New("received nil map")
	}

	outputMap := make(map[string]int, len(inputMap))
	for key, value := range inputMap {
		if _, ok := outputMap[value]; !ok {
			outputMap[value] = key
		} else {
			return nil, errors.New(fmt.Sprintf("value %s will become a duplicate key", value))
		}
	}

	return outputMap, nil
}
