package productservice

import (
	"context"
	"database/sql"
	"log"

	"github.com/imnzr/quiet-leaf-cafe/backend/helper"
	"github.com/imnzr/quiet-leaf-cafe/backend/models"
	productrepository "github.com/imnzr/quiet-leaf-cafe/backend/repository/product_repository"
	productweb "github.com/imnzr/quiet-leaf-cafe/backend/web/product_web"
)

type ProductServiceImpl struct {
	ProductRepository productrepository.ProductRepository
	DB                *sql.DB
}

// FindByAll implements ProductService.
func (service *ProductServiceImpl) FindByAll(ctx context.Context) []productweb.ProductDataResponse {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	products, err := service.ProductRepository.FindByAll(ctx, tx)
	if err != nil {
		log.Printf("error finding all products: %v", err)
		return nil
	}
	var productResponse []productweb.ProductDataResponse
	for _, product := range products {
		productResponse = append(productResponse, productweb.ProductDataResponse{
			Product_id:  product.Product_id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Image:       product.Image,
		})
	}
	return productResponse
}

// FindById implements ProductService.
func (service *ProductServiceImpl) FindById(ctx context.Context, product_id int) productweb.ProductResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, product_id)
	if err != nil {
		log.Printf("error finding product with id: %v", err)
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "Product not found",
		}
	}
	return productweb.ProductResponseHandler{
		Success: true,
		Message: "Product found",
		Data: productweb.ProductDataResponse{
			Product_id:  product.Product_id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Image:       product.Image,
		},
	}
}

// Delete implements ProductService.
func (service *ProductServiceImpl) Delete(ctx context.Context, product_id int) productweb.ProductResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, product_id)
	if err != nil {
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "Product not found",
		}
	}

	err = service.ProductRepository.Delete(ctx, tx, product)
	if err != nil {
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "failed to delete product",
		}
	}

	return productweb.ProductResponseHandler{
		Success: true,
		Message: "Product deleted successfully",
		Data:    nil,
	}
}

// Save implements ProductService.
func (service *ProductServiceImpl) Create(ctx context.Context, request productweb.ProductCreateRequest) productweb.ProductResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Image:       request.Image,
	}

	savedProduct, err := service.ProductRepository.Save(ctx, tx, product)
	if err != nil {
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "Error saving product",
			Data:    nil,
		}
	}
	return productweb.ProductResponseHandler{
		Success: true,
		Message: "Product created successfully",
		Data: productweb.ProductDataResponse{
			Product_id:  savedProduct.Product_id,
			Name:        savedProduct.Name,
			Description: savedProduct.Description,
			Price:       savedProduct.Price,
			Image:       savedProduct.Image,
		},
	}
}

// Search implements ProductService.
func (service *ProductServiceImpl) Search(ctx context.Context, keyword string) ([]productweb.ProductDataResponse, error) {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	products, err := service.ProductRepository.Search(ctx, tx, keyword)
	if err != nil {
		return nil, err
	}

	var responses []productweb.ProductDataResponse

	for _, product := range products {
		responses = append(responses, productweb.ProductDataResponse{
			Product_id:  product.Product_id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Image:       product.Image,
		})
	}
	return responses, nil
}

// UpdateDescription implements ProductService.
func (service *ProductServiceImpl) UpdateDescription(ctx context.Context, request productweb.ProductUpdateDescription) productweb.ProductResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Product_id)
	if err != nil {
		log.Printf("product with id %d not found: %v", request.Product_id, err)
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "Product not found",
			Data:    nil,
		}
	}

	product.Description = request.Description

	_, err = service.ProductRepository.UpdateDescription(ctx, tx, product)
	if err != nil {
		log.Printf("error updating description with id %d: %v", request.Product_id, err)
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "error updating product",
			Data:    nil,
		}
	}
	updatedProductDescription, err := service.ProductRepository.FindById(ctx, tx, request.Product_id)
	if err != nil {
		log.Printf("error fetching updated product with id %d: %v", request.Product_id, err)
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "error fetching updated product",
			Data:    nil,
		}
	}
	return productweb.ProductResponseHandler{
		Success: true,
		Message: "Product description update successfully",
		Data: productweb.ProductDataResponse{
			Name:        updatedProductDescription.Name,
			Description: updatedProductDescription.Description,
			Price:       updatedProductDescription.Price,
			Image:       updatedProductDescription.Image,
		},
	}
}

// UpdateName implements ProductService.
func (service *ProductServiceImpl) UpdateName(ctx context.Context, request productweb.ProductUpdateName) productweb.ProductResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Product_id)
	if err != nil {
		log.Printf("user with id %d not found: %v", request.Product_id, err)
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "Product not found",
			Data:    nil,
		}
	}
	product.Name = request.Name

	_, err = service.ProductRepository.UpdateName(ctx, tx, product)
	if err != nil {
		log.Printf("error updating name with id %d: %v", request.Product_id, err)
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "Error updating product name",
			Data:    nil,
		}
	}
	updatedProductName, err := service.ProductRepository.FindById(ctx, tx, request.Product_id)
	if err != nil {
		log.Printf("error fetching updated product with id %d: %v", request.Product_id, err)
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "error fetching updated product",
			Data:    nil,
		}
	}
	return productweb.ProductResponseHandler{
		Success: true,
		Message: "Product updated name successfully",
		Data: productweb.ProductDataResponse{
			Name:        updatedProductName.Name,
			Description: updatedProductName.Description,
			Price:       updatedProductName.Price,
			Image:       updatedProductName.Image,
		},
	}
}

// UpdatePrice implements ProductService.
func (service *ProductServiceImpl) UpdatePrice(ctx context.Context, request productweb.ProductUpdatePrice) productweb.ProductResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	// mencari produk
	product, err := service.ProductRepository.FindById(ctx, tx, request.Product_id)
	if err != nil {
		log.Printf("product with id %d not found: %v", request.Product_id, err)
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "Product not found",
			Data:    nil,
		}
	}
	product.Price = request.Price

	_, err = service.ProductRepository.UpdatePrice(ctx, tx, product)
	if err != nil {
		log.Printf("error updating price with product id %d: %v", request.Product_id, err)
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "Error fetching update price",
			Data:    nil,
		}
	}

	updatedProductPrice, err := service.ProductRepository.FindById(ctx, tx, request.Product_id)
	if err != nil {
		log.Printf("error fetching update product with id %d: %v", request.Product_id, err)
		return productweb.ProductResponseHandler{
			Success: false,
			Message: "error fetching updated product",
			Data:    nil,
		}
	}

	return productweb.ProductResponseHandler{
		Success: true,
		Message: "Update price product successfully",
		Data: productweb.ProductDataResponse{
			Name:        updatedProductPrice.Name,
			Description: updatedProductPrice.Description,
			Price:       updatedProductPrice.Price,
			Image:       updatedProductPrice.Image,
		},
	}
}

func NewProductService(productRepository productrepository.ProductRepository, db *sql.DB) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                db,
	}
}
