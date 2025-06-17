package productcontroller

import (
	"encoding/json"
	"net/http"

	productservice "github.com/imnzr/quiet-leaf-cafe/backend/service/product_service"
	"github.com/imnzr/quiet-leaf-cafe/backend/utils"
	"github.com/imnzr/quiet-leaf-cafe/backend/web"
	productweb "github.com/imnzr/quiet-leaf-cafe/backend/web/product_web"
	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService productservice.ProductService
}

// Create implements ProductControllerInterface.
func (controller *ProductControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decode := json.NewDecoder(request.Body)
	productCreateRequest := productweb.ProductCreateRequest{}
	err := decode.Decode(&productCreateRequest)
	if err != nil {
		http.Error(writter, "invalid request body", http.StatusBadRequest)
	}

	productResponse := controller.ProductService.Create(request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "Created",
		Data:   productResponse,
	}
	utils.WriteJsonError(writter, http.StatusCreated, webResponse)
}

// Delete implements ProductControllerInterface.
func (controller *ProductControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// Search implements ProductControllerInterface.
func (controller *ProductControllerImpl) Search(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// UpdateDescription implements ProductControllerInterface.
func (controller *ProductControllerImpl) UpdateDescription(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// UpdateName implements ProductControllerInterface.
func (controller *ProductControllerImpl) UpdateName(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// UpdatePrice implements ProductControllerInterface.
func (controller *ProductControllerImpl) UpdatePrice(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func NewProductController(productService productservice.ProductService) ProductControllerInterface {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}
