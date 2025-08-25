package payoutmanagementsystem

type payee struct {
	beneficiaryName string
	beneficiaryCode string
	accNo int
	ifsc string
	bankName string
	email string
	mobile int
	payeeCategory string
}

func NewPayee(name string, code string, accNumber int, ifsc string, bankName string,
	email string, mobile int,payeeCategory string) *payee {
		return &payee{beneficiaryName: name, beneficiaryCode: code, accNo: accNumber, ifsc: ifsc, 
		bankName: bankName, email: email, mobile: mobile, payeeCategory: payeeCategory}
}
