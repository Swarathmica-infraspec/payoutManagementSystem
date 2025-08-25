package payoutmanagementsystem

import (
	"errors"
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

func NewPayee(name string, code string, accNumber int, ifsc string, bankName string,
	email string, mobile int, payeeCategory string) (*payee, error) {
	if numberOfDigits(accNumber) != 10 && numberOfDigits(accNumber) != 16 {
		return nil, errors.New("payoutmanagementsystem: NewPayee: payee should be created with account number of length 10 or 16")
	}
	return &payee{beneficiaryName: name, beneficiaryCode: code, accNo: accNumber, ifsc: ifsc,
		bankName: bankName, email: email, mobile: mobile, payeeCategory: payeeCategory}, nil
}

func numberOfDigits(number int) int {
	numString := strconv.Itoa(number)
	return len(numString)
}
