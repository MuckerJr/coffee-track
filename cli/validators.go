package cli

import (
	"errors"
)

func ValidateNonEmpty(str string) error {
	if str == "" {
		return errors.New("cannot be empty")
	}
	return nil
}
