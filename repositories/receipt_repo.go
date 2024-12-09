package repositories

import (
	"fmt"
	"errors"
	"receipt-processor/models"
	"github.com/google/uuid"
)

var receipts = make(map[string]models.Receipt)

func SaveReceipt(receipt models.Receipt) string {
	id := uuid.New().String()
	receipts[id] = receipt
	return id
}

func GetReceiptById(id string) (models.Receipt, error) {
	receipt, receiptExists := receipts[id]
	if !receiptExists {
		errMsg := fmt.Sprintf("receipt by ID %s not found", id)
		return models.Receipt{}, errors.New(errMsg)
	}
	return receipt, nil
}