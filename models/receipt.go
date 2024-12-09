package models

import (
	"time"
	"math/big"
)

type Receipt struct {
	Id string 
	Retailer string
	PurchaseDate time.Time
	PurchaseTime time.Time
	Total big.Rat
	Items []Item
}