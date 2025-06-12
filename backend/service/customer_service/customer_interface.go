package customerservice

import (
	"context"

	customerweb "github.com/imnzr/quiet-leaf-cafe/backend/web/customer_web"
)

type CustomerService interface {
	Create(ctx context.Context, request customerweb.CustomerCreateRequest) customerweb.CustomerResponseHandler
	Delete(ctx context.Context, customer_id int) customerweb.CustomerResponseHandler
	FindById(ctx context.Context, customer_id int) customerweb.CustomerResponseHandler
	FindByAll(ctx context.Context) []customerweb.CustomerDataResponse
	FindByEmail(ctx context.Context, email string) customerweb.CustomerResponseHandler
	Login(ctx context.Context, request customerweb.CustomerLoginRequest) customerweb.CustomerResponseHandler

	UpdateName(ctx context.Context, request customerweb.CustomerUpdateName) customerweb.CustomerResponseHandler
	UpdateEmail(ctx context.Context, request customerweb.CustomerUpdateEmail) customerweb.CustomerResponseHandler
	UpdatePhoneNumber(ctx context.Context, request customerweb.CustomerUpdatePhoneNumber) customerweb.CustomerResponseHandler
	UpdatePassword(ctx context.Context, request customerweb.CustomerUpdatePassword) customerweb.CustomerResponseHandler
}
