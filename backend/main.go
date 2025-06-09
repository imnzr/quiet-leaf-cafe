package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	customercontroller "github.com/imnzr/quiet-leaf-cafe/backend/controller/customer_controller"
	"github.com/imnzr/quiet-leaf-cafe/backend/database"
	customerrepository "github.com/imnzr/quiet-leaf-cafe/backend/repository/customer_repository"
	customerroutes "github.com/imnzr/quiet-leaf-cafe/backend/routes/customer_routes"
	customerservice "github.com/imnzr/quiet-leaf-cafe/backend/service/customer_service"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := database.GetConnection()
	if err != nil {
		log.Fatal("Failed to connect to database: %w", err)
	}
	defer db.Close()

	fmt.Println("Connected to database successfully")

	customerRepository := customerrepository.NewCustomerRepository()
	customerService := customerservice.NewCustomerService(customerRepository, db)
	customerController := customercontroller.NewCustomerController(customerService)

	router := httprouter.New()
	customerroutes.CustomerRouter(router, customerController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
