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

// UpdateName implements CustomerService.
func (service *CustomerServiceImpl) UpdateName(ctx context.Context, request customerweb.CustomerUpdateName) customerweb.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	// 1. ambil data awal
	customer, err := service.CustomerRepository.FindById(ctx, tx, request.Customer_id)
	if err != nil {
		log.Printf("user with id %d not found: %v", request.Customer_id, err)
		return customerweb.CustomerResponse{}
	}

	// 2. ubah field name
	customer.Name = request.Name

	// 3. update hanya name
	_, err = service.CustomerRepository.UpdateName(ctx, tx, customer)
	if err != nil {
		log.Printf("error updating custome with id %d: %v", request.Customer_id, err)
		return customerweb.CustomerResponse{}
	}

	// 4. ambil ulang customer setelah update
	updatedCustomer, err := service.CustomerRepository.FindById(ctx, tx, request.Customer_id)
	if err != nil {
		log.Printf("error fetching updated customer with id %d: %v", request.Customer_id, err)
		return customerweb.CustomerResponse{}
	}

	// 5. kembalikan data lengkapp
	return customerweb.CustomerResponse{
		Customer_id:  updatedCustomer.Customer_id,
		Name:         updatedCustomer.Name,
		Phone_number: updatedCustomer.Phone_number,
		Email:        updatedCustomer.Email,
	}
}

// UpdateEmail implements CustomerService.
func (service *CustomerServiceImpl) UpdateEmail(ctx context.Context, request customerweb.CustomerUpdateEmail) customerweb.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, request.CustomerId)
	if err != nil {
		log.Printf("user with id %d not found: %v", request.CustomerId, err)
		return customerweb.CustomerResponse{}

	}

	customer.Email = request.Email

	_, err = service.CustomerRepository.UpdateEmail(ctx, tx, customer)
	if err != nil {
		log.Printf("error updating customer with id %d: %v", request.CustomerId, err)
		return customerweb.CustomerResponse{}
	}

	updateCustomer, err := service.CustomerRepository.FindById(ctx, tx, request.CustomerId)
	if err != nil {
		log.Printf("error fetching updated customer with id %d: %v", request.CustomerId, err)
	}

	return customerweb.CustomerResponse{
		Customer_id:  updateCustomer.Customer_id,
		Name:         updateCustomer.Name,
		Phone_number: updateCustomer.Phone_number,
		Email:        updateCustomer.Email,
	}
}

// UpdatePassword implements CustomerService.
func (service *CustomerServiceImpl) UpdatePassword(ctx context.Context, request customerweb.CustomerUpdatePassword) customerweb.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	// 1. ambil data customer berdasarkan ID
	customer, err := service.CustomerRepository.FindById(ctx, tx, request.Customer_Id)
	if err != nil {
		log.Printf("customer with id %d not found: %v", request.Customer_Id, err)
		return customerweb.CustomerResponse{}
	}
	// 2. hash password baru
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error hashing password for customer with id %d: %v", request.Customer_Id, err)
		return customerweb.CustomerResponse{}
	}
	// 3. update password customer
	customer.Password = string(hashedPassword)
	_, err = service.CustomerRepository.UpdatePassword(ctx, tx, customer)
	if err != nil {
		log.Printf("error updating password for customer with id %d: %v", request.Customer_Id, err)
		return customerweb.CustomerResponse{}
	}
	// 4. ambil ulang customer setelah update
	updatedCustomer, err := service.CustomerRepository.FindById(ctx, tx, request.Customer_Id)
	if err != nil {
		log.Printf("error fetching updated customer with id %d: %v", request.Customer_Id, err)
		return customerweb.CustomerResponse{}
	}
	return customerweb.CustomerResponse{
		Customer_id:  updatedCustomer.Customer_id,
		Name:         updatedCustomer.Name,
		Phone_number: updatedCustomer.Phone_number,
		Email:        updatedCustomer.Email,
	}
}

// UpdatePhoneNumber implements CustomerService.
func (service *CustomerServiceImpl) UpdatePhoneNumber(ctx context.Context, request customerweb.CustomerUpdatePhoneNumber) customerweb.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, request.CustomerId)
	if err != nil {
		log.Printf("customer with id %d not found: %v", request.CustomerId, err)
		return customerweb.CustomerResponse{}
	}
	customer.Phone_number = request.Phone_number

	_, err = service.CustomerRepository.UpdatePhoneNumber(ctx, tx, customer)
	if err != nil {
		log.Printf("error updating phone number for customer with id %d: %v", request.CustomerId, err)
	}
	updatedCustomer, err := service.CustomerRepository.FindById(ctx, tx, request.CustomerId)
	if err != nil {
		log.Printf("error fetching updated customer with id %d: %v", request.CustomerId, err)
	}

	return customerweb.CustomerResponse{
		Customer_id:  updatedCustomer.Customer_id,
		Name:         updatedCustomer.Name,
		Phone_number: updatedCustomer.Phone_number,
		Email:        updatedCustomer.Email,
	}
}
