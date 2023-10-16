package customerrepository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kodekage/banking/domain/entities"
	"github.com/kodekage/banking/internal/errors"
	"github.com/kodekage/banking/internal/logger"
)

type Repository interface {
	FindAll(statusFilter string) ([]entities.Customer, *errors.AppError)
	FindById(id string) (*entities.Customer, *errors.AppError)
}

type customerRepository struct {
	sqlClient *sqlx.DB
}

var _ Repository = (*customerRepository)(nil)

func New(dbClient *sqlx.DB) Repository {
	return customerRepository{dbClient}
}

func (c customerRepository) FindAll(statusFilter string) ([]entities.Customer, *errors.AppError) {
	var err error
	customers := make([]entities.Customer, 0)

	if statusFilter != "" {
		selectAllQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = c.sqlClient.Select(&customers, selectAllQuery, statusFilter)
	} else {
		selectAllQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = c.sqlClient.Select(&customers, selectAllQuery)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errors.NewUnexpectedError(err.Error())
	}

	return customers, nil
}

func (c customerRepository) FindById(id string) (*entities.Customer, *errors.AppError) {
	selectByIdQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var customer entities.Customer
	err := c.sqlClient.Get(&customer, selectByIdQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customers " + err.Error())
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
	}

	return &customer, nil
}
