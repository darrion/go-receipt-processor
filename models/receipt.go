package models

type Receipt struct {
	Id string 
	Retailer string
	PurchaseDate string
	PurchaseTime string
	Total string
	Items []Item
}