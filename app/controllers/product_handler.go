package controllers

import (
	"github.com/pranotobudi/gophers-art-backend/app/common"
	"github.com/pranotobudi/gophers-art-backend/app/models"
	"net/http"
)

type ProductResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
	Category    string `json:"category"`
	ImageUrl    string `json:"image"`
}

func ProductsResponseFormatter(products []models.Product) []ProductResponse {
	var formatters []ProductResponse
	for _, product := range products {
		formatter := ProductResponse{
			ID:          product.ID,
			Title:       product.Title,
			Price:       product.Price,
			Rating:      product.Rating,
			Description: product.Description,
			Category:    product.Category,
			ImageUrl:    product.ImageUrl,
		}
		formatters = append(formatters, formatter)
	}
	return formatters
}

type productHandler struct {
	repository models.ProductRepository
}

func NewProductHandler() *productHandler {
	repository := models.NewProductRepository()

	return &productHandler{repository}
}

func (h *productHandler) GetProducts(c App) common.Response {
	// Get Products from repository
	products, err := h.repository.GetProducts()
	if err != nil {
		return common.ResponseErrorFormatter(http.StatusInternalServerError, err)
	}

	// Success ProductResponse
	data := ProductsResponseFormatter(products)

	return common.ResponseFormatter(http.StatusOK, "success", "get products successfull", data)
}
