package handlers

import (
    "net/http"
    "product-api/internal/db"
    "strconv"
"fmt"
    "github.com/gin-gonic/gin"
)

type ProductHandler struct {
    db *db.DB
}

func NewProductHandler(db *db.DB) *ProductHandler {
    return &ProductHandler{db: db}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
    var product db.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.db.CreateProduct(&product); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }
    fmt.Println("Product created successfully:", product)
    c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    product, err := h.db.GetProductByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
    userID := c.Query("user_id")
    priceMin := c.Query("price_min")
    priceMax := c.Query("price_max")
    name := c.Query("name")

    filters := db.ProductFilters{
        UserID:    userID,
        PriceMin:  priceMin,
        PriceMax:  priceMax,
        Name:      name,
    }

    products, err := h.db.ListProducts(filters)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
        return
    }

    c.JSON(http.StatusOK, products)
}
