package payoutmanagementsystem

import (
	"errors"
	"testing"
)

func TestNewPayee(t *testing.T) {
	_, err := NewPayee("abc", "123", 6780003045, "cbin05648", "cbi", "abc@gmail.com", 9123456789, "Employee")
	if err != nil {
		t.Fatal("payee should be created")
	}
}

func TestPayeeCannotBeCreatedWithInvalidAccountNumberOfLengthOtherThan10(t *testing.T) {
	_, err := NewPayee("abc", "123", 678000, "cbin05648", "cbi", "abc@gmail.com", 91234567892, "Employee")
	expectedErr := errors.New("payoutmanagementsystem: NewPayee: payee should be created with account number of length 10 or 16")
	if err == expectedErr {
		t.Fatal("payee should not be created with invalid account number")
	}
}

func TestNewPayeeShouldBeCreatedWithAccountNumberOfLength16(t *testing.T) {
	_, err := NewPayee("abc", "123", 6780002345765432, "cbin05648", "cbi", "abc@gmail.com", 9123456789, "Employee")
	if err != nil {
		t.Fatal("payee should be created with account number of length 16")
	}
}

func TestMobileNumberMustBeOfLength10(t *testing.T) {
	_, err := NewPayee("abc", "123", 6780002345765432, "cbin05648", "cbi", "abc@gmail.com", 912345678, "Employee")
	expectedErr := errors.New("payoutmanagementsystem: NewPayee: length of mobile number must be 10")
	if err == expectedErr {
		t.Fatal("payee should not be created with invalid mobile number")
	}
}

func TestInvalidEmail(t *testing.T) {
	_, err := NewPayee("abc", "123", 6780002345765432, "cbin05648", "cbi", "abc.com", 9123456782, "Employee")
	expectedErr := errors.New("payoutmanagementsystem: NewPayee: email is invalid")
	if err == expectedErr {
		t.Fatal("payee should not be created with invalid email")
	}
}
