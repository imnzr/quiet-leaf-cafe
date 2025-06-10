package customerweb

type CustomerUpdatePassword struct {
	Customer_Id     int
	Password        string
	NewPassword     string
	OldPassword     string
	ConfirmPassword string
}
