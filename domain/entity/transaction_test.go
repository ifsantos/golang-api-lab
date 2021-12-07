package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_IsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 100
	assert.Nil(t, transaction.IsValid())
}
func TestTransaction_IsInvalidWithAmountGreatherThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1001
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for this transaction", err.Error())
}
func TestTransaction_IsInvalidWithAmountLessThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "amount less than 1 not allowed", err.Error())
}
