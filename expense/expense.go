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

func NewExpense(title string, amount float64, dateIncurred string, category string, notes string, payeeID int, receiptURI string) (*expense, error) {
	if title == "" {
		return nil, ErrInvalidTitle
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
