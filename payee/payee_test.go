package payoutmanagementsystem

import (
	"testing"
)

var errTests = []struct {
	beneficiaryName string
	beneficiaryCode string
	accNo           int
	ifsc            string
	bankName        string
	email           string
	mobile          int
	payeeCategory   string
	expectedErr     error
}{
	{"abc", "123", 6780003045, "cbin05648", "cbi", "abc@gmail.com", 9123456789, "Employee", nil},
	{"abc", "123", 678000, "cbin05648", "cbi", "abc@gmail.com", 91234567892, "Employee", ErrInvalidAccountNumber},
	{"abc", "123", 6780002345765432, "cbin05648", "cbi", "abc@gmail.com", 9123456789, "Employee", nil},
	{"abc", "123", 6780002345765432, "cbin05648", "cbi", "abc@gmail.com", 912345678, "Employee", ErrInvalidMobileNumber},
	{"abc", "123", 6780002345765432, "cbin05648", "cbi", "abc.com", 9123456782, "Employee", ErrInvalidEmail},
	{"", "123", 6780002345765432, "cbin05648", "cbi", "abc@gmail.com", 9123456782, "Employee", ErrEmptyName},
	{"abc", "", 6780002345765432, "cbin05648", "cbi", "abc@gmail.com", 9123456782, "Employee", ErrEmptyCode},
}

func TestValidatePayee(t *testing.T) {
	for _, tt := range errTests {
		_, err := NewPayee(tt.beneficiaryName, tt.beneficiaryCode, tt.accNo, tt.ifsc, tt.bankName, tt.email, tt.mobile, tt.payeeCategory)
		if err != tt.expectedErr {
			t.Fatalf("Expected Error: %v but Actual Error: %v", tt.expectedErr, err)
		}
	}
}
