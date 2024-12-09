package main

import (
	"receipt-processor/validators"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
    "receipt-processor/routes"
)

func registerValidators() {
	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validator.RegisterValidation("retailer", validators.ValidateRetailer)
		validator.RegisterValidation("total", validators.ValidateTotal)
		validator.RegisterValidation("short_description", validators.ValidateShortDescription)
		validator.RegisterValidation("price", validators.ValidatePrice)
	}
}

func main() {
    registerValidators()
    router := routes.SetupRouter()
    router.Run(":3000") // Start server on port 8080
}
