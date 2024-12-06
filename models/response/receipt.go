package models

type ReceiptResponse struct {
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total string `json:"total"`
	Items []ItemResponse `json:"items"`
}