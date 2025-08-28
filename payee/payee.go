package payoutmanagementsystem

import (
	"errors"
	"regexp"
	"strconv"
)

type payee struct {
	id              int
	beneficiaryName string
	beneficiaryCode string
	accNo           int
	ifsc            string
	bankName        string
	email           string
	mobile          int
	payeeCategory   string
}

var ErrInvalidAccountNumber = errors.New("payoutmanagementsystem.NewPayee: account number must be of 10 or 16 digits only")
var ErrInvalidEmail = errors.New("payoutmanagementsystem.NewPayee: invalid email format")
var ErrInvalidMobileNumber = errors.New("payoutmanagementsystem.NewPayee: mobile number must be of 10 digits only")
var ErrEmptyName = errors.New("payoutmanagementsystem.NewPayee: name should not be empty")
var ErrEmptyCode = errors.New("payoutmanagementsystem.NewPayee: code should not be empty")
var ErrInvalidIFSC = errors.New("payoutmanagementsystem.NewPayee: invalid ifsc code")
var ErrInvalidBankName = errors.New("payoutmanagementsystem.NewPayee: invalid bank name")

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
	if !checkIFSC(ifsc) {
		return nil, ErrInvalidIFSC
	}
	if len(bankName) > 50 {
		return nil, ErrInvalidBankName
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

func checkIFSC(ifsc string) bool {
	if len(ifsc) != 10 {
		return false
	}
	matchAlphaForFirstFourChars, _ := regexp.MatchString("([A-Z]+)", ifsc[:4])
	if !matchAlphaForFirstFourChars {
		return false
	}
	if ifsc[4] != '0' {
		return false
	}
	// TODO: check last 5 digits of ifsc
	// fmt.Println(ifsc[5:10])
	// LastFiveCharacters := regexp.MustCompile(ifsc[5:10])
	// matchNumForLastFiveChars, _ := regexp.MatchString("[1-9]+", LastFiveCharacters)
	// if !matchNumForLastFiveChars {
	// 	return false
	// }
	return true
}
