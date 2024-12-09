package services

import (
	"unicode"
	"math/big"
)

func EvenlyDivides(dividend int, divisor int) bool {
	return dividend % divisor == 0
}

func IsAlphaNumChar(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsNumber(char)
}

func IsRoundDollar(amount *big.Rat) bool {
	return amount.IsInt()
}

func IsMultipleOfQuarter(amount *big.Rat) bool {
	quotient := new(big.Rat).Mul(amount, big.NewRat(4, 1))
	return quotient.IsInt()
}
