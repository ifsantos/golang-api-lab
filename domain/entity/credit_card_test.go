package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("1234123412341234", "john silver", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5515318818029095", "MasterCard User", 12, 2024, 123)
	assert.Nil(t, err)

}
func TestCreditExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5515318818029095", "john silver", 13, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5515318818029095", "john silver", 0, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())
}

func TestCreditExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)

	_, err := NewCreditCard("5515318818029095", "john silver", 12, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())
}
