package productapi

import (
	productcontroller "github.com/imnzr/quiet-leaf-cafe/backend/controller/product_controller"
	"github.com/julienschmidt/httprouter"
)

func ProductRouter(router *httprouter.Router, productController productcontroller.ProductControllerInterface) {
	router.POST("/api/product/create", productController.Create)
}
