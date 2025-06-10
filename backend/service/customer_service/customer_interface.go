package customerservice

import (
	"context"

	customerweb "github.com/imnzr/quiet-leaf-cafe/backend/web/customer_web"
)

type CustomerService interface {
	Create(ctx context.Context, request customerweb.CustomerCreateRequest) customerweb.CustomerResponse
	Delete(ctx context.Context, customer_id int)
	FindById(ctx context.Context, customer_id int) customerweb.CustomerResponse
	FindByAll(ctx context.Context) []customerweb.CustomerResponse
	FindByEmail(ctx context.Context, email string) (customerweb.CustomerResponse, error)
	Login(ctx context.Context, request customerweb.CustomerLoginRequest) (customerweb.CustomerResponse, error)

	UpdateName(ctx context.Context, request customerweb.CustomerUpdateName) customerweb.CustomerResponse
	UpdateEmail(ctx context.Context, request customerweb.CustomerUpdateEmail) customerweb.CustomerResponse
	UpdatePhoneNumber(ctx context.Context, request customerweb.CustomerUpdatePhoneNumber) customerweb.CustomerResponse
	UpdatePassword(ctx context.Context, request customerweb.CustomerUpdatePassword) customerweb.CustomerResponse
}
