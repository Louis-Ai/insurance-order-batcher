package utility

import (
	"fmt"
	"regexp"
)

func ValidatePostcode(postcode string) error {
	if len(postcode) > 8 {
		return fmt.Errorf("postcode exceeds maximum of 9 characters")
	}

	r := regexp.MustCompile("^[a-zA-Z0-9 ]+$")

	if !r.MatchString(postcode) {
		return fmt.Errorf("postcode contains invalid characters (only number letters and spaces allowed)")
	}
	return nil
}
