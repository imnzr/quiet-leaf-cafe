package paymentservice

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/xendit/xendit-go/v7"
	"github.com/xendit/xendit-go/v7/invoice"
)

type PaymentServiceImpl struct{}

func NewPaymentService() PaymentServerInterface {
	return &PaymentServiceImpl{}
}

// CreatePayment implements PaymentServerInterface.
func (p *PaymentServiceImpl) CreatePayment(orderId int64, customerName string, orderNumber string, totalAmount float64) (string, error) {

	apiKey := os.Getenv("XENDIT_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("xendit key environment variable is not set")
	}
	// inisialiasi xendit client
	xenditClient := xendit.NewClient(apiKey)

	ExternalId := "ORD-123456789"
	Email := "hello@customer.com"
	Customer := "customer"
	Description := "Pembayaran Quiet Leaf Cafe"

	// buat invoice request
	createInvoiceRequest := *invoice.NewCreateInvoiceRequest(orderNumber, totalAmount)
	createInvoiceRequest.Customer = &invoice.CustomerObject{
		CustomerId: *invoice.NewNullableString(&ExternalId),
		Email:      *invoice.NewNullableString(&Email),
		GivenNames: *invoice.NewNullableString(&Customer),
	}
	createInvoiceRequest.Description = &Description

	// jalankan api call
	resp, r, err := xenditClient.InvoiceApi.CreateInvoice(context.Background()).
		CreateInvoiceRequest(createInvoiceRequest).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.CreateInvoice`: %v\n", err.Error())
		b, _ := json.Marshal(err.FullError())
		fmt.Fprintf(os.Stderr, "full error struct: %v\n", string(b))
		fmt.Fprintf(os.Stderr, "full http response: %v\n", r)
	}
	return resp.InvoiceUrl, nil
}
