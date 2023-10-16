package accountcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodekage/banking/dto"
	"github.com/kodekage/banking/service/accountservice"
	"github.com/kodekage/banking/utils"
)

func Mount(r *mux.Router, s accountservice.Service) {
	c := accountController{service: s}

	r.HandleFunc("/accounts", c.CreateNewAccount).Methods(http.MethodPost)
}

type accountController struct {
	service accountservice.Service
}

func (a accountController) CreateNewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateAccountRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		account, appErr := a.service.CreateNewAccount(request)

		if appErr != nil {
			utils.WriteResponse(w, appErr.Code, appErr.Message)
		} else {
			utils.WriteResponse(w, http.StatusCreated, account)
		}
	}
}
