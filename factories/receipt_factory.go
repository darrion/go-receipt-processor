package factories

import (
	"receipt-processor/models"
	"receipt-processor/responses"
)

func ToReceiptResponse(receipt models.Receipt) responses.ReceiptResponse {
    itemResponses := make([]responses.ItemResponse, len(receipt.Items))
    for index, item := range receipt.Items {
        itemResponses[index] = responses.ItemResponse{
            ShortDescription: item.ShortDescription,
            Price:            item.Price,
        }
    }

    return responses.ReceiptResponse{
        Retailer:     receipt.Retailer,
        PurchaseDate: receipt.PurchaseDate,
        PurchaseTime: receipt.PurchaseTime,
        Total:        receipt.Total,
        Items:        itemResponses,
    }
}

func ToReceipt(receiptResponse responses.ReceiptResponse, id string) models.Receipt {
    items := make([]models.Item, len(receiptResponse.Items))
    for index, itemResponse := range receiptResponse.Items {
        items[index] = models.Item{
            ShortDescription: itemResponse.ShortDescription,
            Price:            itemResponse.Price,
        }
    }

    return models.Receipt{
        Id:           id,
        Retailer:     receiptResponse.Retailer,
        PurchaseDate: receiptResponse.PurchaseDate,
        PurchaseTime: receiptResponse.PurchaseTime,
        Total:        receiptResponse.Total,
        Items:        items,
    }
}
