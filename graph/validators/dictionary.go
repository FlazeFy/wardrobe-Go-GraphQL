package validators

import (
	"fmt"
	"strings"
)

func CreateDictionaryInput(dictionaryName string, dictionaryType string) error {
	dictionaryName = strings.TrimSpace(dictionaryName)
	dictionaryType = strings.TrimSpace(dictionaryType)

	// Dictionary Name
	if dictionaryName == "" {
		return fmt.Errorf("dictionaryName is required")
	}
	if len(dictionaryName) < 3 {
		return fmt.Errorf("dictionaryName min length is 3 characters")
	}
	if len(dictionaryName) > 75 {
		return fmt.Errorf("dictionaryName max length is 75 characters")
	}

	// Dictionary Type
	if dictionaryType == "" {
		return fmt.Errorf("dictionaryType is required")
	}
	if len(dictionaryType) > 36 {
		return fmt.Errorf("dictionaryType max length is 36 characters")
	}

	return nil
}
