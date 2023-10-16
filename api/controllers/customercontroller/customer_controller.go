package customercontroller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodekage/banking/dto"
	"github.com/kodekage/banking/service/customerservice"
	"github.com/kodekage/banking/utils"
)

func Mount(r *mux.Router, s customerservice.Service) {
	c := customerController{service: s}

	r.HandleFunc("/customers", c.getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id:[0-9]+}", c.getCustomer).Methods(http.MethodGet)
}

type customerController struct {
	service customerservice.Service
}

func (c *customerController) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	customers, err := c.service.GetAllCustomers(dto.CustomerFilters{Status: status})

	if err != nil {
		utils.WriteResponse(w, err.Code, err.AsMessage())
	} else {
		utils.WriteResponse(w, http.StatusOK, customers)
	}
}

func (c *customerController) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := c.service.GetCustomer(id)

	if err != nil {
		utils.WriteResponse(w, err.Code, err.AsMessage())
	} else {
		utils.WriteResponse(w, http.StatusOK, customer)

	}
}
