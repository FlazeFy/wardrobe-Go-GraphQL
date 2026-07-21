package validators

import (
	"fmt"
	"strings"
)

func CreateQuestionInput(question string) error {
	trimmed := strings.TrimSpace(question)

	if trimmed == "" {
		return fmt.Errorf("question is required")
	}
	if len(trimmed) < 10 {
		return fmt.Errorf("question min length is 10 characters")
	}
	if len(trimmed) > 500 {
		return fmt.Errorf("question max length is 500 characters")
	}

	return nil
}
