package accountrepository

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/kodekage/banking/domain/entities"
	"github.com/kodekage/banking/internal/errors"
	"github.com/kodekage/banking/internal/logger"
)

type Repository interface {
	Save(a entities.Account) (*entities.Account, *errors.AppError)
}

type accountRepository struct {
	sqlClient *sqlx.DB
}

var _ Repository = (*accountRepository)(nil)

func New(dbClient *sqlx.DB) Repository {
	return accountRepository{dbClient}
}

func (ar accountRepository) Save(a entities.Account) (*entities.Account, *errors.AppError) {
	insertQuery := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

	result, err := ar.sqlClient.Exec(insertQuery, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error " + err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while while getting last insert id for new account " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error " + err.Error())
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}
