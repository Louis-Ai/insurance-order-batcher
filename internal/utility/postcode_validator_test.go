package utility

import (
	"testing"
)

func TestPostcodeValidation(t *testing.T) {

	tests := []struct {
		testname string
		postcode string
		expected bool
	}{
		{"Valid postcode", "B99 4SR", true},
		{"Invalid postcode", "B22-3QX", false},
	}

	for _, tc := range tests {
		t.Run(tc.testname, func(t *testing.T) {
			err := ValidatePostcode(tc.postcode)

			if (err != nil) != tc.expected {
				t.Errorf("Postcode: %s error: %v wantErr: %v", tc.postcode, err, tc.expected)
			}
		})
	}

}
