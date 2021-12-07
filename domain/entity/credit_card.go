package entity

import (
	"errors"
	"regexp"
)

type CreditCard struct {
	number          string
	name            string
	expirationMonth int
	expirationYear  int
	cvv             int
}

func NewCreditCard(number string, name string, expirationMonth int, expirantionYear int, cvv int) (*CreditCard, error) {

	cc := &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirationYear:  expirantionYear,
		cvv:             cvv,
	}
	err := cc.IsValid()
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *CreditCard) IsValid() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14})$`)

	if !re.MatchString(c.number) {
		return errors.New("invalid credit card number")
	}
	return nil
}
