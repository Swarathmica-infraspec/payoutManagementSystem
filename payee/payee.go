package payoutmanagementsystem

import (
	"errors"
	"regexp"
	"strconv"
)

type payee struct {
	beneficiaryName string
	beneficiaryCode string
	accNo           int
	ifsc            string
	bankName        string
	email           string
	mobile          int
	payeeCategory   string
}

var ErrInvalidAccountNumber = errors.New("payoutmanagementsystem: NewPayee: payee should be created with account number of length 10 or 16")
var ErrInvalidEmail = errors.New("payoutmanagementsystem: NewPayee: email is invalid")
var ErrInvalidMobileNumber = errors.New("payoutmanagementsystem: NewPayee: length of mobile number must be 10")
var ErrEmptyName = errors.New("payoutmanagementsystem: NewPayee: payee should not be created with empty name")
var ErrEmptyCode = errors.New("payoutmanagementsystem: NewPayee: payee should not be created with empty code")

func NewPayee(name string, code string, accNumber int, ifsc string, bankName string,
	email string, mobile int, payeeCategory string) (*payee, error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	if code == "" {
		return nil, ErrEmptyCode
	}
	if numberOfDigits(accNumber) != 10 && numberOfDigits(accNumber) != 16 {
		return nil, ErrInvalidAccountNumber
	}
	if numberOfDigits(mobile) != 10 {
		return nil, ErrInvalidMobileNumber
	}
	if !checkEmailFormat(email) {
		return nil, ErrInvalidEmail
	}
	return &payee{beneficiaryName: name, beneficiaryCode: code, accNo: accNumber, ifsc: ifsc,
		bankName: bankName, email: email, mobile: mobile, payeeCategory: payeeCategory}, nil
}

func numberOfDigits(number int) int {
	numString := strconv.Itoa(number)
	return len(numString)
}

// Is this check enough for email? domain name has to be checked if it exists?
func checkEmailFormat(email string) bool {
	match, _ := regexp.MatchString("([a-z]+)(@)([a-z]+)(.)[com]", email)
	return match
}
