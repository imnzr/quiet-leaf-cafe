package models

type OrderItems struct {
	Order_detail_id int
	Order_id        int
	Product_id      int
	Temperature_id  int
	Cupsize_id      int
	Sweetness_id    int
	Topping_id      int
	AddOn_id        int
	Quantity        int
	// Price           int
}

type OrderRequest struct {
	Customer_id int
	Items       []OrderItems
}
