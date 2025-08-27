package payoutmanagementsystem

import (
	"errors"
)

type expense struct {
	title        string
	amount       float64
	dateIncurred string
	category     string
	notes        string
	payeeID      int
	receiptURI   string
}

var ErrInvalidTitle = errors.New("payoutmanagementsystem.NewExpense: title should not be empty")
var ErrInvalidAmount = errors.New("payoutmanagementsystem.NewExpense: amount must be greater than 0")

func NewExpense(title string, amount float64, dateIncurred string, category string, notes string, payeeID int, receiptURI string) (*expense, error) {
	if title == "" {
		return nil, ErrInvalidTitle
	}
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	return &expense{
		title:        title,
		amount:       amount,
		dateIncurred: dateIncurred,
		category:     category,
		notes:        notes,
		payeeID:      payeeID,
		receiptURI:   receiptURI,
	}, nil
}
