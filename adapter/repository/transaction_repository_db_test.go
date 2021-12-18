package repository

import (
	"os"
	"testing"

	"github.com/ifsantos/golang-api-lab/adapter/repository/fixture"
	"github.com/ifsantos/golang-api-lab/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	transactionRepositoryDb := NewTransactionRepositoryDb(db)
	err := transactionRepositoryDb.Insert("1", "1", 12.1, entity.APPROVED, "")
	assert.Nil(t, err)
}
