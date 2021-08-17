package utils

import (
	"fmt"
	"os"
	"time"
)

func WriteSomeEntries(entriesCount int, filepath string) {
	writeEntry := func(filepath string) error {
		file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}(file)
		entry := fmt.Sprintf("new entry at %v\n", time.Now())
		_, _ = file.WriteString(entry)
		return nil
	}

	for i := 0; i < entriesCount; i++ {
		err := writeEntry(filepath)
		if err != nil {
			panic(err)
		}
	}
}
