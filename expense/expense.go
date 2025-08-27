package payoutmanagementsystem

import (
	"errors"
	"regexp"
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
var ErrInvalidDate = errors.New("payoutmanagementsystem.NewExpense: invalid date values or format (YYYY-MM-DD)")

func NewExpense(title string, amount float64, dateIncurred string, category string, notes string, payeeID int, receiptURI string) (*expense, error) {
	if title == "" {
		return nil, ErrInvalidTitle
	}
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	if !checkDateFormat(dateIncurred) {
		return nil, ErrInvalidDate
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

func checkDateFormat(date string) bool {
	pattern := `^(202[5-9]|20(3\d|4\d|50))-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`
	match, _ := regexp.MatchString(pattern, date)
	return match
}
