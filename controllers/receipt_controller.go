package controllers

import (
    "net/http"
    "receipt-processor/requests"
    "receipt-processor/services"
    "github.com/gin-gonic/gin"
)

func ProcessReceipt(c *gin.Context) {
    var receiptRequest requests.ReceiptRequest
    if err := c.ShouldBindJSON(&receiptRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid"})
        return
    }
    id := services.SaveReceipt(receiptRequest)
    c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetReceiptById(c *gin.Context) {
	id := c.Param("id")
	receipt, err := services.GetReceiptById(id)
	if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
        return
    }
	c.JSON(http.StatusOK, gin.H{"receipt": receipt})
}

func GetPoints(c *gin.Context) {
	id := c.Param("id")
	points, err := services.GetPoints(id)
	if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
        return
    }
	c.JSON(http.StatusOK, gin.H{"points": points})
}	
