package productroutes

import (
	productcontroller "github.com/imnzr/quiet-leaf-cafe/backend/controller/product_controller"
	"github.com/julienschmidt/httprouter"
)

func ProductRouter(router *httprouter.Router, productController productcontroller.ProductControllerInterface) {
	router.POST("/api/product/create", productController.Create)
	router.GET("/product/search/", productController.Search)
	router.GET("/api/products", productController.FindByAll)
	router.GET("/api/product/:productId", productController.FindById)
	router.DELETE("/product/delete/:productId", productController.Delete)

	router.PUT("/product/update/description/:productId", productController.UpdateName)
	router.PUT("/product/update/name/:productId", productController.UpdateDescription)
	router.PUT("/product/update/price/:productId", productController.UpdatePrice)
}
