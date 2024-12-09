package main

import (
    
    "item-api/handlers"
    "item-api/internal/database"
    "github.com/gin-gonic/gin"
"fmt"
)
func main() {
    database := database.NewDB()

    router := gin.Default()

    itemHandler := handlers.NewProductHandler(database)

    // Root endpoint
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to the Product API",
        })
    })

    // Log for debugging
    fmt.Println("Registering /items endpoint")
    router.POST("/items", itemHandler.CreateProduct)
    router.GET("/items/:id", itemHandler.GetProductByID)
    router.GET("/items", itemHandler.ListProducts)

    // Start server
    fmt.Println("Starting server on port 8080...")
    router.Run(":8080")
}


