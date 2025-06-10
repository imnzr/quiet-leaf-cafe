package customercontroller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CustomerControllerInterface interface {
	Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writter http.ResponseWriter, request *http.Request, params httprouter.Params)

	UpdateName(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateEmail(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdatePhoneNumber(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdatePassword(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}
