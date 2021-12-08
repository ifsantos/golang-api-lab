package entity

import "errors"

const (
	REJECTED = "rejected"
	APPROVED = "approved"
)

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("you dont have limit for this transaction")
	}
	if t.Amount < 1 {
		return errors.New("amount less than 1 not allowed")
	}
	return nil
}
