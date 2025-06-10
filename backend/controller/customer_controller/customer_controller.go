package customercontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	customerservice "github.com/imnzr/quiet-leaf-cafe/backend/service/customer_service"
	"github.com/imnzr/quiet-leaf-cafe/backend/utils"
	"github.com/imnzr/quiet-leaf-cafe/backend/web"
	customerweb "github.com/imnzr/quiet-leaf-cafe/backend/web/customer_web"
	"github.com/julienschmidt/httprouter"
)

type CustomerControllerImpl struct {
	CustomerService customerservice.CustomerService
}

func NewCustomerController(customerService customerservice.CustomerService) CustomerControllerInterface {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

// Create implements CustomerControllerInterface.
func (controller *CustomerControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decode := json.NewDecoder(request.Body)
	customerCreateRequest := customerweb.CustomerCreateRequest{}
	err := decode.Decode(&customerCreateRequest)
	if err != nil {
		http.Error(writter, "Invalid request body", http.StatusBadRequest)
	}

	customerResponse := controller.CustomerService.Create(request.Context(), customerCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   customerResponse,
	}
	utils.WriteJsonError(writter, http.StatusCreated, webResponse)
}

// Delete implements CustomerControllerInterface.
func (controller *CustomerControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	if err != nil {
		http.Error(writter, "Invalid customer ID", http.StatusBadRequest)
		return
	}
	controller.CustomerService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "Accepted",
	}
	utils.WriteJsonError(writter, http.StatusAccepted, webResponse)
}

// FindByAll implements CustomerControllerInterface.
func (controller *CustomerControllerImpl) FindByAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customer := controller.CustomerService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customer,
	}
	utils.WriteJsonError(writter, http.StatusOK, webResponse)
}

// FindById implements CustomerControllerInterface.
func (controller *CustomerControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	if err != nil {
		http.Error(writter, "Invalid customer ID", http.StatusBadRequest)
		return
	}
	customer := controller.CustomerService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customer,
	}
	utils.WriteJsonError(writter, http.StatusOK, webResponse)
}

// Login implements CustomerControllerInterface.
func (controller *CustomerControllerImpl) Login(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	customerLoginRequest := customerweb.CustomerLoginRequest{}
	err := decoder.Decode(&customerLoginRequest)
	if err != nil {
		http.Error(writter, "Invalid request body", http.StatusBadRequest)
		return
	}

	customerResponse, err := controller.CustomerService.Login(request.Context(), customerLoginRequest)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusUnauthorized)
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	}
	utils.WriteJsonError(writter, http.StatusOK, webResponse)

}

// UpdateName implements CustomerControllerInterface.
func (controller *CustomerControllerImpl) UpdateName(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	customerUpdateName := customerweb.CustomerUpdateName{}
	err := decoder.Decode(&customerUpdateName)
	if err != nil {
		http.Error(writter, "invalid request body", http.StatusBadRequest)
	}

	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	if err != nil {
		http.Error(writter, "invalid customer ID", http.StatusBadRequest)
		return
	}

	customerUpdateName.Customer_id = id

	customerResponse := controller.CustomerService.UpdateName(request.Context(), customerUpdateName)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	}
	utils.WriteJsonError(writter, http.StatusOK, webResponse)
}

// UpdateEmail implements CustomerControllerInterface.
func (controller *CustomerControllerImpl) UpdateEmail(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	customerUpdateEmail := customerweb.CustomerUpdateEmail{}
	err := decoder.Decode(&customerUpdateEmail)
	if err != nil {
		http.Error(writter, "invalid request body", http.StatusBadRequest)
	}

	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	if err != nil {
		http.Error(writter, "invalid customer ID", http.StatusBadRequest)
	}

	customerUpdateEmail.CustomerId = id

	customerResponse := controller.CustomerService.UpdateEmail(request.Context(), customerUpdateEmail)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	}
	utils.WriteJsonError(writter, http.StatusOK, webResponse)
}

// UpdatePassword implements CustomerControllerInterface.
func (controller *CustomerControllerImpl) UpdatePassword(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	customerUpdatePassword := customerweb.CustomerUpdatePassword{}
	err := decoder.Decode(&customerUpdatePassword)
	if err != nil {
		http.Error(writter, "invalid request body", http.StatusBadRequest)
	}

	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	if err != nil {
		http.Error(writter, "invalid customer ID", http.StatusBadRequest)
	}

	customerUpdatePassword.Customer_Id = id

	customerResponse := controller.CustomerService.UpdatePassword(request.Context(), customerUpdatePassword)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	}
	utils.WriteJsonError(writter, http.StatusOK, webResponse)
}

// UpdatePhoneNumber implements CustomerControllerInterface.
func (controller *CustomerControllerImpl) UpdatePhoneNumber(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	customerUpdatePhone := customerweb.CustomerUpdatePhoneNumber{}
	err := decoder.Decode(&customerUpdatePhone)
	if err != nil {
		http.Error(writter, "invalid request body", http.StatusBadRequest)
	}

	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	if err != nil {
		http.Error(writter, "invalid customer ID", http.StatusBadRequest)
	}

	customerUpdatePhone.CustomerId = id

	customerResponse := controller.CustomerService.UpdatePhoneNumber(request.Context(), customerUpdatePhone)

	WebResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	}
	utils.WriteJsonError(writter, http.StatusOK, WebResponse)
}
