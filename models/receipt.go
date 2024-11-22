package models

type Receipt struct {
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	Total string `json:"total"`
	Items []Item `json:"items"`
}