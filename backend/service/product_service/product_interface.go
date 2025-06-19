package productservice

import (
	"context"

	productweb "github.com/imnzr/quiet-leaf-cafe/backend/web/product_web"
)

type ProductService interface {
	Create(ctx context.Context, request productweb.ProductCreateRequest) productweb.ProductResponseHandler
	Delete(ctx context.Context, product_id int) productweb.ProductResponseHandler
	Search(ctx context.Context, keyword string) ([]productweb.ProductDataResponse, error)
	UpdateDescription(ctx context.Context, request productweb.ProductUpdateDescription) productweb.ProductResponseHandler
	UpdateName(ctx context.Context, request productweb.ProductUpdateName) productweb.ProductResponseHandler
	UpdatePrice(ctx context.Context, request productweb.ProductUpdatePrice) productweb.ProductResponseHandler
	FindById(ctx context.Context, product_id int) productweb.ProductResponseHandler
	FindByAll(ctx context.Context) []productweb.ProductDataResponse
}
