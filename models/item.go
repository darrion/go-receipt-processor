package models

import (
	"math/big"
)

type Item struct {
	ShortDescription string
	Price big.Rat
}