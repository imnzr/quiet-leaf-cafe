package customerroutes

import (
	customercontroller "github.com/imnzr/quiet-leaf-cafe/backend/controller/customer_controller"
	"github.com/julienschmidt/httprouter"
)

func CustomerRouter(router *httprouter.Router, customerController customercontroller.CustomerControllerInterface) {
	router.GET("/api/customers", customerController.FindByAll)
	router.GET("/api/customer/:customerId", customerController.FindById)
	router.POST("/api/customer/register", customerController.Create)
	router.POST("/api/customer/login", customerController.Login)
	router.DELETE("/api/customer/:customerId", customerController.Delete)
	router.PUT("/api/update/:customerId/name", customerController.UpdateName)
}
