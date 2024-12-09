package validators

import (
	"regexp"
	"github.com/go-playground/validator/v10"
)

func ValidateRetailer(fl validator.FieldLevel) bool {
	return ValidateRegex(fl, RETAILER_PATTERN)
}

func ValidateTotal(fl validator.FieldLevel) bool {
	return ValidateRegex(fl, TOTAL_PATTERN)
}

func ValidateShortDescription(fl validator.FieldLevel) bool {
	return ValidateRegex(fl, SHORT_DESCRIPTION_PATTERN)
}

func ValidatePrice(fl validator.FieldLevel) bool {
	return ValidateRegex(fl, PRICE_PATTERN)
}

func ValidateRegex(fl validator.FieldLevel, pattern string) bool {
	re, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}
	
	return re.MatchString(fl.Field().String())
}