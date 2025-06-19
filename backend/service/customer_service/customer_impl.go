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
func (service *CustomerServiceImpl) FindByEmail(ctx context.Context, email string) customerweb.CustomerResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customer, err := service.CustomerRepository.FindByEmail(ctx, tx, email)
	if err != nil {
		log.Printf("error finding customer with email %s: %v", email, err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Customer not found",
		}
	}
	return customerweb.CustomerResponseHandler{
		Success: true,
		Message: "Customer found",
		Data: customerweb.CustomerDataResponse{
			CustomerId:  customer.Customer_id,
			Name:        customer.Name,
			PhoneNumber: customer.Phone_number,
			Email:       customer.Email,
		},
	}
}

// Delete implements CustomerService.
func (service *CustomerServiceImpl) Delete(ctx context.Context, customer_id int) customerweb.CustomerResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, customer_id)
	if err != nil {
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Customer not found",
		}
	}
	service.CustomerRepository.Delete(ctx, tx, customer)

	return customerweb.CustomerResponseHandler{
		Success: true,
		Message: "Customer deleted successfully",
		Data:    nil,
	}
}

// FindByAll implements CustomerService.
func (service *CustomerServiceImpl) FindByAll(ctx context.Context) []customerweb.CustomerDataResponse {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customers, err := service.CustomerRepository.FindByAll(ctx, tx)
	if err != nil {
		log.Printf("error finding all customers: %v", err)
		return nil
	}
	var customerResponse []customerweb.CustomerDataResponse
	for _, customer := range customers {
		customerResponse = append(customerResponse, customerweb.CustomerDataResponse{
			CustomerId:  customer.Customer_id,
			Name:        customer.Name,
			PhoneNumber: customer.Phone_number,
			Email:       customer.Email,
		})
	}
	return customerResponse
}

// FindById implements CustomerService.
func (service *CustomerServiceImpl) FindById(ctx context.Context, customer_id int) customerweb.CustomerResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, customer_id)
	if err != nil {
		log.Printf("error finding customer with id: %v", err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Customer not found",
		}
	}
	return customerweb.CustomerResponseHandler{
		Success: true,
		Message: "Customer found",
		Data: customerweb.CustomerDataResponse{
			CustomerId:  customer.Customer_id,
			Name:        customer.Name,
			PhoneNumber: customer.Phone_number,
			Email:       customer.Email,
		},
	}

}

// Login implements CustomerService.
func (service *CustomerServiceImpl) Login(ctx context.Context, request customerweb.CustomerLoginRequest) (customerweb.CustomerResponseHandler, error) {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customer, err := service.CustomerRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Customer not found",
		}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password))
	if err != nil {
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Invalid password",
		}, err
	}
	return customerweb.CustomerResponseHandler{
		Success: true,
		Message: "Login successful",
		Data: customerweb.CustomerDataResponse{
			CustomerId:  customer.Customer_id,
			Name:        customer.Name,
			PhoneNumber: customer.Phone_number,
			Email:       customer.Email,
		},
	}, nil

}

// Save implements CustomerService.
func (service *CustomerServiceImpl) Create(ctx context.Context, request customerweb.CustomerCreateRequest) customerweb.CustomerResponseHandler {
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
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Error saving customer",
			Data:    nil,
		}
	}
	return customerweb.CustomerResponseHandler{
		Success: true,
		Message: "Customer created successfully",
		Data: customerweb.CustomerDataResponse{
			CustomerId:  savedUser.Customer_id,
			Name:        savedUser.Name,
			PhoneNumber: savedUser.Phone_number,
			Email:       savedUser.Email,
		},
	}
}

// UpdateName implements CustomerService.
func (service *CustomerServiceImpl) UpdateName(ctx context.Context, request customerweb.CustomerUpdateName) customerweb.CustomerResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	// 1. ambil data awal
	customer, err := service.CustomerRepository.FindById(ctx, tx, request.Customer_id)
	if err != nil {
		log.Printf("user with id %d not found: %v", request.Customer_id, err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Customer not found",
			Data:    nil,
		}
	}
	// 2. ubah field name
	customer.Name = request.Name

	// 3. update hanya name
	_, err = service.CustomerRepository.UpdateName(ctx, tx, customer)
	if err != nil {
		log.Printf("error updating customer with id %d: %v", request.Customer_id, err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Error updating customer",
			Data:    nil,
		}
	}

	// 4. ambil ulang customer setelah update
	updatedCustomer, err := service.CustomerRepository.FindById(ctx, tx, request.Customer_id)
	if err != nil {
		log.Printf("error fetching updated customer with id %d: %v", request.Customer_id, err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Error fetching updated customer",
			Data:    nil,
		}
	}

	// 5. kembalikan data name
	return customerweb.CustomerResponseHandler{
		Success: true,
		Message: "Customer updated name successfully",
		Data: customerweb.CustomerDataResponse{
			Name: updatedCustomer.Name,
		},
	}
}

// UpdateEmail implements CustomerService.
func (service *CustomerServiceImpl) UpdateEmail(ctx context.Context, request customerweb.CustomerUpdateEmail) customerweb.CustomerResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, request.CustomerId)
	if err != nil {
		log.Printf("user with id %d not found: %v", request.CustomerId, err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Customer not found",
			Data:    nil,
		}

	}

	customer.Email = request.Email

	_, err = service.CustomerRepository.UpdateEmail(ctx, tx, customer)
	if err != nil {
		log.Printf("error updating customer with id %d: %v", request.CustomerId, err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Error updating customer",
		}
	}

	_, err = service.CustomerRepository.FindById(ctx, tx, request.CustomerId)
	if err != nil {
		log.Printf("error fetching updated customer with id %d: %v", request.CustomerId, err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Error fetching updated customer",
			Data:    nil,
		}
	}

	return customerweb.CustomerResponseHandler{
		Success: true,
		Message: "Customer update email successfully",
		Data:    nil,
	}
}

// UpdatePassword implements CustomerService.
func (service *CustomerServiceImpl) UpdatePassword(ctx context.Context, request customerweb.CustomerUpdatePassword) customerweb.CustomerResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	// 1. ambil data customer berdasarkan ID
	customer, err := service.CustomerRepository.FindById(ctx, tx, request.Customer_Id)
	if err != nil {
		log.Printf("customer with id %d not found: %v", request.Customer_Id, err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Customer not found",
		}
	}
	// 2. hash password baru
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error hashing password for customer with id %d: %v", request.Customer_Id, err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Error hashing password",
		}
	}
	// 3. update password customer
	customer.Password = string(hashedPassword)
	_, err = service.CustomerRepository.UpdatePassword(ctx, tx, customer)
	if err != nil {
		log.Printf("error updating password for customer with id %d: %v", request.Customer_Id, err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Error updating password",
		}
	}
	// 4. ambil ulang customer setelah update
	_, err = service.CustomerRepository.FindById(ctx, tx, request.Customer_Id)
	if err != nil {
		log.Printf("error fetching updated customer with id %d: %v", request.Customer_Id, err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Error fetching updated customer",
		}
	}
	return customerweb.CustomerResponseHandler{
		Success: true,
		Message: "Password updated successfully",
	}
}

// UpdatePhoneNumber implements CustomerService.
func (service *CustomerServiceImpl) UpdatePhoneNumber(ctx context.Context, request customerweb.CustomerUpdatePhoneNumber) customerweb.CustomerResponseHandler {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, request.CustomerId)
	if err != nil {
		log.Printf("customer with id %d not found: %v", request.CustomerId, err)
		return customerweb.CustomerResponseHandler{
			Success: false,
			Message: "Customer not found",
			Data:    nil,
		}
	}
	customer.Phone_number = request.Phone_number

	_, err = service.CustomerRepository.UpdatePhoneNumber(ctx, tx, customer)
	if err != nil {
		log.Printf("error updating phone number for customer with id %d: %v", request.CustomerId, err)
	}
	_, err = service.CustomerRepository.FindById(ctx, tx, request.CustomerId)
	if err != nil {
		log.Printf("error fetching updated customer with id %d: %v", request.CustomerId, err)
	}

	return customerweb.CustomerResponseHandler{
		Success: true,
		Message: "Customer updated phone number successfully",
		Data:    nil,
	}
}
