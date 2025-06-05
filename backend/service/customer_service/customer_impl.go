package customerservice

import (
	"context"
	"database/sql"
	"log"

	"github.com/imnzr/quiet-leaf-cafe/backend/helper"
	"github.com/imnzr/quiet-leaf-cafe/backend/models"
	customerrepository "github.com/imnzr/quiet-leaf-cafe/backend/repository/customer_repository"
	customerweb "github.com/imnzr/quiet-leaf-cafe/backend/web/customer_web"
	"golang.org/x/crypto/bcrypt"
)

type CustomerServiceImpl struct {
	CustomerRepository customerrepository.CustomerRepository
	DB                 *sql.DB
}

func NewCustomerService(customerRepository customerrepository.CustomerRepository, db *sql.DB) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		DB:                 db,
	}
}

// FindByEmail implements CustomerService.
func (service *CustomerServiceImpl) FindByEmail(ctx context.Context, email string) (customerweb.CustomerResponse, error) {
	panic("unimplemented")
}

// Delete implements CustomerService.
func (service *CustomerServiceImpl) Delete(ctx context.Context, customer_id int) {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, customer_id)
	if err != nil {
		log.Printf("error finding customer with id:  %v", err)
	}
	service.CustomerRepository.Delete(ctx, tx, customer)
}

// FindByAll implements CustomerService.
func (service *CustomerServiceImpl) FindByAll(ctx context.Context) []customerweb.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customers, err := service.CustomerRepository.FindByAll(ctx, tx)
	if err != nil {
		log.Printf("error finding all customers: %v", err)
		return nil
	}
	var customerResponse []customerweb.CustomerResponse
	for _, customer := range customers {
		customerResponse = append(customerResponse, customerweb.CustomerResponse{
			Customer_id:  customer.Customer_id,
			Name:         customer.Name,
			Email:        customer.Email,
			Phone_number: customer.Phone_number,
		})
	}
	return customerResponse
}

// FindById implements CustomerService.
func (service *CustomerServiceImpl) FindById(ctx context.Context, customer_id int) customerweb.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, customer_id)
	if err != nil {
		log.Printf("error finding customer with id: %v", err)
	}

	customerResponse := customerweb.CustomerResponse{
		Customer_id:  customer.Customer_id,
		Name:         customer.Name,
		Email:        customer.Email,
		Phone_number: customer.Phone_number,
	}
	return customerResponse

}

// Login implements CustomerService.
func (service *CustomerServiceImpl) Login(ctx context.Context, request customerweb.CustomerLoginRequest) (customerweb.CustomerResponse, error) {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customer, err := service.CustomerRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		log.Printf("customer with email %s not found: %v", request.Email, err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password))
	if err != nil {
		log.Printf("error comparing password for customer with email %s: %v", request.Email, err)
	}

	return customerweb.CustomerResponse{
		Customer_id:  customer.Customer_id,
		Name:         customer.Name,
		Email:        customer.Email,
		Phone_number: customer.Phone_number,
	}, nil
}

// Save implements CustomerService.
func (service *CustomerServiceImpl) Create(ctx context.Context, request customerweb.CustomerCreateRequest) customerweb.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error hashing password:", err)
	}
	customer := models.Customer{
		Name:         request.Name,
		Email:        request.Email,
		Phone_number: request.Phone_number,
		Password:     string(hashedPassword),
	}
	savedUser, err := service.CustomerRepository.Save(ctx, tx, customer)
	if err != nil {
		log.Printf("error saving customer: %v", err)
	}
	return customerweb.CustomerResponse{
		Customer_id:  savedUser.Customer_id,
		Name:         savedUser.Name,
		Email:        savedUser.Email,
		Phone_number: savedUser.Phone_number,
	}
}

// Update implements CustomerService.
func (service *CustomerServiceImpl) Update(ctx context.Context, request customerweb.CustomerUpdateRequest) customerweb.CustomerResponse {
	panic("unimplemented")
}
