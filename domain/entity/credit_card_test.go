package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T){
	creditCard, err := NewCreditCard("1234123412341234", "john silver", 12, 2024, 123)
	assert.Equal(t, expected: "Invalid credit card number", err.Error())

}
