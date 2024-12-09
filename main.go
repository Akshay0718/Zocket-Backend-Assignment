package main

import (
    
    "product-api/handlers"
    "product-api/internal/db"
    "github.com/gin-gonic/gin"
"fmt"
)
func main() {
    database := db.NewDB()

    router := gin.Default()

    productHandler := handlers.NewProductHandler(database)

    // Root endpoint
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to the Product API",
        })
    })

    // Log for debugging
    fmt.Println("Registering /products endpoint")
    router.POST("/products", productHandler.CreateProduct)
    router.GET("/products/:id", productHandler.GetProductByID)
    router.GET("/products", productHandler.ListProducts)

    // Start server
    fmt.Println("Starting server on port 8080...")
    router.Run(":8080")
}


