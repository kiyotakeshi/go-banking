package app

import (
	"banking/dto"
	"banking/errs"
	mock_service "banking/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	// arrange
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockService := mock_service.NewMockCustomerService(controller)
	dummyCustomers := []dto.CustomerResponse{
		{"100", "mike", "tokyo", "001001", "1998-01-01", "1"},
		{"101", "popcorn", "osaka", "001002", "1988-11-01", "1"},
	}
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
	customerHandler := CustomerHandlers{mockService}

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.getCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// assertion
	if recorder.Code != http.StatusOK {
		t.Error("failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	// arrange
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockService := mock_service.NewMockCustomerService(controller)
	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedError("merror"))
	customerHandler := CustomerHandlers{mockService}

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.getCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// assertion
	if recorder.Code != http.StatusInternalServerError {
		t.Error("failed while testing the status code")
	}
}
