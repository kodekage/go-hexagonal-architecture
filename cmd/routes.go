package cmd

import (
	"github.com/gorilla/mux"
	"github.com/kodekage/banking/api/controllers/accountcontroller"
	"github.com/kodekage/banking/api/controllers/customercontroller"
	"github.com/kodekage/banking/repositories/accountrepository"
	"github.com/kodekage/banking/repositories/customerrepository"
	"github.com/kodekage/banking/service/accountservice"
	"github.com/kodekage/banking/service/customerservice"
	"github.com/kodekage/banking/utils"
)

func setupRoutes() *mux.Router {
	router := mux.NewRouter()
	sqlClient := utils.SqlClient()

	// define routes
	{
		customerRepo := customerrepository.New(sqlClient)
		customerService := customerservice.New(customerRepo)
		customercontroller.Mount(router, customerService)
	}

	{
		accountRepo := accountrepository.New(sqlClient)
		accountService := accountservice.New(accountRepo)
		accountcontroller.Mount(router, accountService)
	}

	return router
}
