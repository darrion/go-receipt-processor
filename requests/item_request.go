package requests

type ItemRequests struct {
	ShortDescription string `json:"shortDescription" binding:"required,short_description"`
	Price string `json:"price" binding:"required,price"`
}