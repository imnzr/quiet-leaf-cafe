package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	customercontroller "github.com/imnzr/quiet-leaf-cafe/backend/controller/customer_controller"
	productcontroller "github.com/imnzr/quiet-leaf-cafe/backend/controller/product_controller"
	"github.com/imnzr/quiet-leaf-cafe/backend/database"
	customerrepository "github.com/imnzr/quiet-leaf-cafe/backend/repository/customer_repository"
	productrepository "github.com/imnzr/quiet-leaf-cafe/backend/repository/product_repository"
	customerroutes "github.com/imnzr/quiet-leaf-cafe/backend/routes/customer_routes"
	productroutes "github.com/imnzr/quiet-leaf-cafe/backend/routes/product_routes"
	customerservice "github.com/imnzr/quiet-leaf-cafe/backend/service/customer_service"
	productservice "github.com/imnzr/quiet-leaf-cafe/backend/service/product_service"
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

	productRepository := productrepository.NewProductRepository()
	productService := productservice.NewProductService(productRepository, db)
	productController := productcontroller.NewProductController(productService)

	router := httprouter.New()
	customerroutes.CustomerRouter(router, customerController)
	productroutes.ProductRouter(router, productController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
