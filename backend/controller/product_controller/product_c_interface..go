package productcontroller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerInterface interface {
	Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	Search(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateDescription(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateName(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdatePrice(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}
