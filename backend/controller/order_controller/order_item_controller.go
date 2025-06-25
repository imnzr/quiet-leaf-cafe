package ordercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/imnzr/quiet-leaf-cafe/backend/models"
	orderservice "github.com/imnzr/quiet-leaf-cafe/backend/service/order_service"
	"github.com/imnzr/quiet-leaf-cafe/backend/utils"
	"github.com/imnzr/quiet-leaf-cafe/backend/web"
	"github.com/julienschmidt/httprouter"
)

type OrderItemController struct {
	OrderService orderservice.OrderService
}

// CreateOrder implements OrderControllerInterface.
func (controller *OrderItemController) CreateOrder(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decode := json.NewDecoder(request.Body)
	orderItemRequest := models.OrderRequest{}
	err := decode.Decode(&orderItemRequest)
	if err != nil {
		http.Error(writter, "invalid request body", http.StatusBadRequest)
	}

	paymentUrl, err := controller.OrderService.CreateOrder(request.Context(), orderItemRequest)
	if err != nil {
		http.Error(writter, "invalid to create order", http.StatusBadRequest)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   paymentUrl,
	}
	utils.WriteJsonSuccess(writter, http.StatusOK, webResponse)
}

func NewOrderControll(orderService orderservice.OrderService) OrderControllerInterface {
	return &OrderItemController{
		OrderService: orderService,
	}
}
