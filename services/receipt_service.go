package services

import (
	"errors"
	"receipt-processor/models"
	"receipt-processor/models/response"
	"receipt-processor/repositories"
	"receipt-processor/factories"
	"github.com/google/uuid"
)

func SaveReceipt(receipt models.Receipt) (string, error) {
	id := uuid.New().String()
	if err := repositories.SaveReceipt(id, receipt); err != nil {
		return "", err
	}
	return id, nil
}

func GetReceiptById(id string) (response.ReceiptResponse, error) {
	receipt, error := repositories.GetReceiptById(id)
	receiptResponse := factories.ToReceiptResponse(receipt)
	return receiptResponse, error
}