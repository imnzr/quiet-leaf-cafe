package models

type OrderItems struct {
	Order_detail_id int
	Product_id      int
	Temperature_id  int
	Cupsize_id      int
	Sweetness_id    int
	Topping_id      int
	AddOn_id        int
	Quantity        int
}

type OrderRequest struct {
	Customer_id int
	Items       []OrderItems
}

type OrderWithPaymentResponse struct {
	Order_id    int64
	Payment_url string
}

type Order struct {
	Order_id      int
	Order_number  string
	Status        string
	TotalAmount   int64
	CreatedAt     string
	Customer_name string
}
