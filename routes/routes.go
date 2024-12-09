package routes

import (
    "github.com/gin-gonic/gin"
    "receipt-processor/controllers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    api := router.Group("/receipts")
    {
        api.POST("/process", controllers.ProcessReceipt)
		api.GET("/:id", controllers.GetReceiptById)
        api.GET("/:id/points", controllers.GetPoints)
    }

    return router
}
