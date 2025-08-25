package payoutmanagementsystem

import (
	"testing"
)


func TestNewPayee(t *testing.T) {
	payee := NewPayee("abc","123",6780003045,"cbin05648","cbi", "abc@gmail.com", 91234567892, "Employee")
	if payee == nil {
		t.Fatal("payee should be created")
	}
}
