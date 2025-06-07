package customercontroller

import (
	"net/http"

	customerservice "github.com/imnzr/quiet-leaf-cafe/backend/service/customer_service"
	"github.com/julienschmidt/httprouter"
)

type CustomerControllerImpl struct {
	CustomerService customerservice.CustomerService
}

// Create implements CustomerControllerInterface.
func (c *CustomerControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// Delete implements CustomerControllerInterface.
func (c *CustomerControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// FindByAll implements CustomerControllerInterface.
func (c *CustomerControllerImpl) FindByAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// FindById implements CustomerControllerInterface.
func (c *CustomerControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// Login implements CustomerControllerInterface.
func (c *CustomerControllerImpl) Login(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// Update implements CustomerControllerInterface.
func (c *CustomerControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func NewCustomerControllre(customerService customerservice.CustomerService) CustomerControllerInterface {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}
