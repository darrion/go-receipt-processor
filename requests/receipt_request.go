package requests

type ReceiptRequest struct {
	Retailer string `json:"retailer" binding:"required,retailer"`
	PurchaseDate string `json:"purchaseDate" binding:"required,datetime=2006-01-02"`
	PurchaseTime string `json:"purchaseTime" binding:"required,datetime=15:04"`
	Total string `json:"total" binding:"required,total"`
	Items []ItemRequests `json:"items" binding:"required,dive"`
}