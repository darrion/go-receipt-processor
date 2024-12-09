package services

import (
	"receipt-processor/requests"
	"receipt-processor/responses"
	"receipt-processor/repositories"
	"receipt-processor/factories"
)

func SaveReceipt(receiptRequest requests.ReceiptRequest) string {
	receipt := factories.RequestToReceipt(receiptRequest)
	id := repositories.SaveReceipt(receipt)
	return id
}

func GetReceiptById(id string) (responses.ReceiptResponse, error) {
	receipt, err := repositories.GetReceiptById(id)
	receiptResponse := factories.ReceiptToResponse(id, receipt)
	return receiptResponse, err
}

func GetPoints(id string) (int, error) {
	receipt, err := repositories.GetReceiptById(id)

	if err != nil {
		return 0, err
	}

	pointsForRetailer := GetPointsForRetailer(receipt.Retailer)
	pointsForTotal := GetPointsForTotal(&receipt.Total)
	pointsForEveryTwoItems := GetPointsForEveryTwoItems(receipt.Items)
	pointsForDescriptions := GetPointsForItemDescriptions(receipt.Items)
	pointsForPurchaseDate := GetPointsForPurchaseDate(receipt.PurchaseDate)
	pointsForPurchaseTime := GetPointsForPurchaseTime(receipt.PurchaseTime)
	
	points := pointsForRetailer
	points += pointsForTotal
	points += pointsForEveryTwoItems
	points += pointsForDescriptions
	points += pointsForPurchaseDate
	points += pointsForPurchaseTime

	return points, nil
}