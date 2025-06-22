package orderroutes

import (
	ordercontroller "github.com/imnzr/quiet-leaf-cafe/backend/controller/order_controller"
	"github.com/julienschmidt/httprouter"
)

func OrderRouter(router *httprouter.Router, orderController ordercontroller.OrderControllerInterface) {
	router.POST("/api/product/order", orderController.CreateOrder)
}
