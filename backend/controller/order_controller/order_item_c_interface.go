package ordercontroller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OrderControllerInterface interface {
	CreateOrder(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}
