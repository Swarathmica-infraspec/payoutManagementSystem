package payoutmanagementsystem

import (
	"testing"
)

var errTests = []struct {
	testName        string
	errMsg          string
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
	{"Validate Account Number Length", "invalid account number: 9 digits is used", "abc", "123", 678000234, "CBIN056489", "cbi", "abc@gmail.com", 9123456789, "Employee", ErrInvalidAccountNumber},
	{"Validate Account Number Length", "invalid account number: 11 digits is used", "abc", "123", 67800023445, "CBIN056489", "cbi", "abc@gmail.com", 9123456789, "Employee", ErrInvalidAccountNumber},
	{"Validate Account Number Length", "invalid account number: 15 digits is used", "abc", "123", 678000234576543, "CBIN056489", "cbi", "abc@gmail.com", 9123456789, "Employee", ErrInvalidAccountNumber},
	{"Validate Account Number Length", "invalid account number: 17 digits is used", "abc", "123", 67800023457654324, "CBIN056489", "cbi", "abc@gmail.com", 9123456789, "Employee", ErrInvalidAccountNumber},
	{"Validate Mobile Number Length", "invalid mobile number: 9 digits is used", "abc", "123", 6780002345765432, "CBIN056489", "cbi", "abc@gmail.com", 912345678, "Employee", ErrInvalidMobileNumber},
	{"Validate Mobile Number Length", "invalid mobile number: 11 digits is used", "abc", "123", 6780002345765432, "CBIN056489", "cbi", "abc@gmail.com", 91234567891, "Employee", ErrInvalidMobileNumber},
	{"Validate Email", "invalid email: @ symbol is not used", "abc", "123", 6780002345765432, "CBIN056489", "cbi", "abc.com", 9123456782, "Employee", ErrInvalidEmail},
	{"Validate Email", "invalid name: name is empty", "", "123", 6780002345765432, "CBIN056489", "cbi", "abc@gmail.com", 9123456782, "Employee", ErrEmptyName},
	{"Validate Name", "invalid code: code is empty", "abc", "", 6780002345765432, "CBIN056489", "cbi", "abc@gmail.com", 9123456782, "Employee", ErrEmptyCode},
	{"Validate IFSC", "invalid ifsc: the numerals are missed", "abc", "123", 6700345678, "CBIN0789", "cbi", "abc@gmail.com", 9123456666, "Employee", ErrInvalidIFSC},
	{"Validate IFSC", "invalid ifsc: the alphabets are small", "abc", "123", 6700345678, "cbin045667", "cbi", "abc@gmail.com", 9123456666, "Employee", ErrInvalidIFSC},
	{"Validate IFSC", "invalid ifsc: the alphabets is used as part of branch code", "abc", "123", 6700345678, "CBIN0456ab", "cbi", "abc@gmail.com", 9123456666, "Employee", ErrInvalidIFSC},
	{"Validate Bank Name", "invalid bank name: bank name exceeds 50 characters", "abc", "123", 6700345678, "CBIN045667", "cbicbicbicbicbicbicbicbicbicbicbicbicbicbicbicbicbicbicbicbicbicbicbicbicbi", "abc@gmail.com", 9123456666, "Employee", ErrInvalidBankName},
}

func TestValidatePayee(t *testing.T) {
	for _, tt := range errTests {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := NewPayee(tt.beneficiaryName, tt.beneficiaryCode, tt.accNo, tt.ifsc, tt.bankName, tt.email, tt.mobile, tt.payeeCategory)
			if err != tt.expectedErr {
				t.Fatalf("Error Test Case: %v , Expected Error: %v but Actual Error: %v", tt.testName, tt.expectedErr, err)
			}
		})
	}
}

var validPayeeFieldsTests = []struct {
	beneficiaryName string
	beneficiaryCode string
	accNo           int
	ifsc            string
	bankName        string
	email           string
	mobile          int
	payeeCategory   string
}{
	{"abc", "123", 6780003045, "CBIN056489", "cbi", "abc@gmail.com", 9123456789, "Employee"},
	{"abc", "123", 6780002345765432, "CBIN056489", "cbi", "abc@gmail.com", 9123456789, "Employee"},
}

func TestPayeeCanBeCreatedWithValidValues(t *testing.T) {
	for _, tt := range validPayeeFieldsTests {
		_, err := NewPayee(tt.beneficiaryName, tt.beneficiaryCode, tt.accNo, tt.ifsc, tt.bankName, tt.email, tt.mobile, tt.payeeCategory)
		if err != nil {
			t.Fatalf("Payee should be created but got error: %v", err)
		}
	}
}
