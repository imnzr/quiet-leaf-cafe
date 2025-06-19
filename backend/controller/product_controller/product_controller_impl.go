package productcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	productservice "github.com/imnzr/quiet-leaf-cafe/backend/service/product_service"
	"github.com/imnzr/quiet-leaf-cafe/backend/utils"
	"github.com/imnzr/quiet-leaf-cafe/backend/web"
	productweb "github.com/imnzr/quiet-leaf-cafe/backend/web/product_web"
	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService productservice.ProductService
}

// FindByAll implements ProductControllerInterface.
func (controller *ProductControllerImpl) FindByAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	product := controller.ProductService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}
	utils.WriteJsonError(writter, http.StatusOK, webResponse)
}

// FindById implements ProductControllerInterface.
func (controller *ProductControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(writter, "invalid product id", http.StatusBadRequest)
		return
	}
	product := controller.ProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}
	utils.WriteJsonError(writter, http.StatusOK, webResponse)
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
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(writter, "Invalid product ID", http.StatusBadRequest)
		return
	}

	controller.ProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "Accepted",
	}
	utils.WriteJsonError(writter, http.StatusAccepted, webResponse)
}

// Search implements ProductControllerInterface.
func (controller *ProductControllerImpl) Search(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	keyword := request.URL.Query().Get("q")
	if keyword == "" {
		http.Error(writter, "query parameter 'q' is required", http.StatusBadRequest)
	}

	result, err := controller.ProductService.Search(request.Context(), keyword)
	if err != nil {
		http.Error(writter, "failed to search product", http.StatusInternalServerError)
		return
	}

	writter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writter).Encode(result)
}

// UpdateDescription implements ProductControllerInterface.
func (controller *ProductControllerImpl) UpdateDescription(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	productUpdateDescription := productweb.ProductUpdateDescription{}
	err := decoder.Decode(&productUpdateDescription)
	if err != nil {
		http.Error(writter, "invalid request body", http.StatusBadRequest)
	}

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(writter, "invalid customer Id", http.StatusBadRequest)
		return
	}

	productUpdateDescription.Product_id = id

	productResponse := controller.ProductService.UpdateDescription(request.Context(), productUpdateDescription)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}
	utils.WriteJsonError(writter, http.StatusOK, webResponse)
}

// UpdateName implements ProductControllerInterface.
func (controller *ProductControllerImpl) UpdateName(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	productUpdateName := productweb.ProductUpdateName{}
	err := decoder.Decode(&productUpdateName)
	if err != nil {
		http.Error(writter, "invalid", http.StatusBadRequest)
	}

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(writter, "invalid product id", http.StatusBadRequest)
		return
	}

	productUpdateName.Product_id = id

	productResponse := controller.ProductService.UpdateName(request.Context(), productUpdateName)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}
	utils.WriteJsonError(writter, http.StatusOK, webResponse)
}

// UpdatePrice implements ProductControllerInterface.
func (controller *ProductControllerImpl) UpdatePrice(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	productUpdatePrice := productweb.ProductUpdatePrice{}
	err := decoder.Decode(&productUpdatePrice)
	if err != nil {
		http.Error(writter, "invalid request body", http.StatusBadRequest)
	}

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(writter, "invalid product id", http.StatusBadRequest)
	}

	productUpdatePrice.Product_id = id

	productResponse := controller.ProductService.UpdatePrice(request.Context(), productUpdatePrice)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}
	utils.WriteJsonError(writter, http.StatusOK, webResponse)
}

func NewProductController(productService productservice.ProductService) ProductControllerInterface {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}
