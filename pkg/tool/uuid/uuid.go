package uuid

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

func uuidValidate(uuidString *string, invalid chan string) {
	_, err := uuid.Parse(*uuidString)
	if err != nil {
		invalid <- *uuidString + ": " + err.Error() + "; "
	} else {
		invalid <- ""
	}
}

// ConcurrencyValidate - validates a slice of uuid strings
func ConcurrencyValidate(stringSlice *[]string) error {
	var invalidString string
	invalidUUIDs := make(chan string)
	for _, i := range *stringSlice {
		go uuidValidate(&i, invalidUUIDs)
		invalidString += <-invalidUUIDs
	}
	if invalidString != "" {
		invalidString = strings.TrimSuffix(invalidString, "; ")
		return errors.New(invalidString)
	}
	return nil
}

// Validate - validates uuid
func Validate(uuidString string) error {
	if _, err := uuid.Parse(uuidString); err != nil {
		return err
	}
	return nil
}