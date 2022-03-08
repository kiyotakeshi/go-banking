package app

import (
	"banking/dto"
	"banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId

		transaction, appErr := h.service.MakeTransaction(request)

		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, transaction)
		}
	}
}