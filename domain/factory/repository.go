package factory

import "github.com/ifsantos/golang-api-lab/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
