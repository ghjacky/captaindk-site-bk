package utils

import "fmt"

func WrapError(err error, message string) error {
	if err != nil {
		return fmt.Errorf("%s: %s", message, err.Error())
	} else {
		return nil
	}
}
