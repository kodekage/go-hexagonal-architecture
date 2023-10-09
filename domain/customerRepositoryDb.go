package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kodekage/banking/errs"
	"github.com/kodekage/banking/logger"
)

type CustomerRepositoryDb struct {
	sqlClient *sqlx.DB
}

func (c CustomerRepositoryDb) FindAll(filters Filters) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	if filters.Status != "" {
		selectAllQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = c.sqlClient.Select(&customers, selectAllQuery, filters.Status)
	} else {
		selectAllQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = c.sqlClient.Select(&customers, selectAllQuery)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return customers, nil
}

func (c CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	selectByIdQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var customer Customer
	err := c.sqlClient.Get(&customer, selectByIdQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	sqlClient, err := sqlx.Open("mysql", "root:rootpw@/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	sqlClient.SetConnMaxLifetime(time.Minute * 3)
	sqlClient.SetMaxOpenConns(10)
	sqlClient.SetMaxIdleConns(10)

	return CustomerRepositoryDb{sqlClient}
}
