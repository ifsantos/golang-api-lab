package entity

import (
	"errors"
	"regexp"
	"time"
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
	err := c.validateNumber()
	if err != nil {
		return err
	}
	err = c.validateMonth()
	if err != nil {
		return err
	}
	err = c.validateYear()
	if err != nil {
		return err
	}
	return nil
}
func (c *CreditCard) validateNumber() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14})$`)

	if !re.MatchString(c.number) {
		return errors.New("invalid credit card number")
	}
	return nil
}

func (c *CreditCard) validateMonth() error {
	if c.expirationMonth > 0 && c.expirationMonth < 13 {
		return nil
	}
	return errors.New("invalid expiration month")
}
func (c *CreditCard) validateYear() error {
	if c.expirationYear >= time.Now().Year() {
		return nil
	}
	return errors.New("invalid expiration year")
}
