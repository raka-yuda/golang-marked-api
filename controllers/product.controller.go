package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Product struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ListProductsParams struct {
	Limit  int32
	Offset int32
}

type ProductController struct {
	// db  *db.Queries
	ctx context.Context
}

func NewProductController(ctx context.Context) *ProductController {
	return &ProductController{ctx}
}

// Dummy function to simulate retrieving products
func (pc *ProductController) listProductsDummy(args *ListProductsParams) ([]Product, error) {
	products := []Product{
		{ID: uuid.New(), Name: "Product 1"},
		{ID: uuid.New(), Name: "Product 2"},
		{ID: uuid.New(), Name: "Product 3"},
	}

	start := args.Offset
	end := start + int32(args.Limit)
	if start >= int32(len(products)) {
		return []Product{}, nil
	}
	if end > int32(len(products)) {
		end = int32(len(products))
	}

	return products[start:end], nil
}

// Retrieve all records handlers
func (pc *ProductController) GetAllProducts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageID, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)
	offset := (reqPageID - 1) * reqLimit

	args := &ListProductsParams{
		Limit:  int32(reqLimit),
		Offset: int32(offset),
	}

	products, err := pc.listProductsDummy(args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed to retrieve products", "error": err.Error()})
		return
	}

	if products == nil {
		products = []Product{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   "Sucessfully retrieved all products",
		"size":     len(products),
		"products": products},
	)
}
