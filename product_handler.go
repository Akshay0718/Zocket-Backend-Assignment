package handlers

import (
    "net/http"
    "product-api/internal/db"
    "strconv"
"fmt"
    "github.com/gin-gonic/gin"
)

type ProductHandler struct {
    databaseConnector *db.DB
}

func NewProductHandler(databaseConnector *db.DB) *ProductHandler {
    return &ProductHandler{databaseConnector: databaseConnector}
}

func (handler *ProductHandler) CreateProduct(c *gin.Context) {
    var item db.Product
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := handler.databaseConnector.CreateProduct(&item); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }
    fmt.Println("Product created successfully:", item)
    c.JSON(http.StatusCreated, item)
}

func (handler *ProductHandler) GetProductByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
        return
    }

    item, err := handler.databaseConnector.GetProductByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, item)
}

func (handler *ProductHandler) ListProducts(c *gin.Context) {
    userIdentifier := c.Query("user_id")
    minPrice := c.Query("price_min")
    maxPrice := c.Query("price_max")
    itemName := c.Query("name")

    queryFilters := db.ProductFilters{
        UserID:    userIdentifier,
        PriceMin:  minPrice,
        PriceMax:  maxPrice,
        Name:      itemName,
    }

    products, err := handler.databaseConnector.ListProducts(queryFilters)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
        return
    }

    c.JSON(http.StatusOK, products)
}
