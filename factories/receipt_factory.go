package factories

import (
    "fmt"
	"receipt-processor/models"
	"receipt-processor/responses"
    "receipt-processor/requests"
)

func ReceiptToResponse(id string, receipt models.Receipt) responses.ReceiptResponse {
    itemResponses := make([]responses.ItemResponse, len(receipt.Items))
    for index, item := range receipt.Items {
        itemResponses[index] = responses.ItemResponse{
            ShortDescription: item.ShortDescription,
            Price:            ConvertBigRatToStrDollarAmt(&item.Price),
        }
    }

    return responses.ReceiptResponse{
        Id:           id,
        Retailer:     receipt.Retailer,
        PurchaseDate: ConvertDateObjToStr(&receipt.PurchaseDate),
        PurchaseTime: ConvertTimeObjToStr(&receipt.PurchaseTime),
        Total:        ConvertBigRatToStrDollarAmt(&receipt.Total),
        Items:        itemResponses,
    }
}

func RequestToReceipt(receiptRequest requests.ReceiptRequest) models.Receipt {
    items := make([]models.Item, len(receiptRequest.Items))
    for index, itemResponse := range receiptRequest.Items {
        items[index] = models.Item{
            ShortDescription: itemResponse.ShortDescription,
            Price:            *ConvertStrNumToRatNum(&itemResponse.Price),
        }
    }

    purchaseDate, err := ConvertDateStrToObj(&receiptRequest.PurchaseDate)

    if err != nil {
        fmt.Println("Error in ToReceipt factory at line 38")
    }

    return models.Receipt{
        Retailer:     receiptRequest.Retailer,
        PurchaseDate: purchaseDate,
        PurchaseTime: ConvertTimeStrToObj(&receiptRequest.PurchaseTime),
        Total:        *ConvertStrNumToRatNum(&receiptRequest.Total),
        Items:        items,
    }
}

func ResponseToReceipt(receiptResponse responses.ReceiptResponse, id string) models.Receipt {
    items := make([]models.Item, len(receiptResponse.Items))
    for index, itemResponse := range receiptResponse.Items {
        items[index] = models.Item{
            ShortDescription: itemResponse.ShortDescription,
            Price:            *ConvertStrNumToRatNum(&itemResponse.Price),
        }
    }

    purchaseDate, err := ConvertDateStrToObj(&receiptResponse.PurchaseDate)

    if err != nil {
        fmt.Println("Error in ToReceipt factory at line 38")
    }

    return models.Receipt{
        Id:           id,
        Retailer:     receiptResponse.Retailer,
        PurchaseDate: purchaseDate,
        PurchaseTime: ConvertTimeStrToObj(&receiptResponse.PurchaseTime),
        Total:        *ConvertStrNumToRatNum(&receiptResponse.Total),
        Items:        items,
    }
}
