package repositories

import (
	"fmt"
	"errors"
	"receipt-processor/models"
)

var receipts = make(map[string]models.Receipt)

func SaveReceipt(id string, receipt models.Receipt) error {
	receipts[id] = receipt
	return nil
}

func GetReceiptById(id string) (models.Receipt, error) {
	receipt, receiptExists := receipts[id]
	if !receiptExists {
		errMsg := fmt.Sprintf("Receipt by ID %s not found.", id)
		return models.Receipt{}, errors.New(errMsg)
	}
	return receipt, nil
}