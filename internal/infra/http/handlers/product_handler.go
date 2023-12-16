package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/tomazcx/products-api/internal/data/factory"
	"github.com/tomazcx/products-api/internal/dto"
	entityPkg "github.com/tomazcx/products-api/pkg/entity"
	"github.com/tomazcx/products-api/pkg/httperr"
)

type ProductHandler struct {
	factory factory.ProductFactory
}

func NewProductHandler(factory factory.ProductFactory) *ProductHandler {
	return &ProductHandler{factory: factory}
}

// Get Many Products godoc
// @Summary      Get many products with pagination
// @Description  Send the desired page, limit and sort operation to list the products.
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        page   query   int  true  "Page"
// @Param        limit   query   int  true  "Limit"
// @Param        sort   query   string  true  "Sort"
// @Success      200 {object} []entity.Product
// @Failure      400
// @Failure      500
// @Router       /products [get]
// @Security BearerAuth
func (h *ProductHandler) GetManyProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Error: invalid page", http.StatusBadRequest)
		return
	}

	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		http.Error(w, "Error: invalid limit", http.StatusBadRequest)
		return
	}

	if sort == "" {
		http.Error(w, "Error: invalid sort option", http.StatusBadRequest)
		return
	}

	showManyProductsUseCase := h.factory.ShowManyProductsUseCase()
	products, err := showManyProductsUseCase.Execute(pageInt, limitInt, sort)

	if err != nil {
		if httpErr, ok := err.(*httperr.HttpError); ok {
			http.Error(w, httpErr.Message, httpErr.StatusCode)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Get Product godoc
// @Summary      Get a product data based on the ID
// @Description  Send the product ID to retrieve all its data.
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id   path   string  true  "ID"
// @Success      200 {object} entity.Product
// @Failure      404
// @Failure      400
// @Failure      500
// @Router       /products/{id} [get]
// @Security BearerAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Error: invalid ID", http.StatusBadRequest)
		return
	}

	parsedID, err := entityPkg.ParseID(id)

	if err != nil {
		http.Error(w, "Error: invalid ID", http.StatusBadRequest)
		return
	}

	showProductUseCase := h.factory.ShowProductUseCase()
	product, err := showProductUseCase.Execute(parsedID.String())

	if err != nil {
		if httpErr, ok := err.(*httperr.HttpError); ok {
			http.Error(w, httpErr.Message, httpErr.StatusCode)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Create Product godoc
// @Summary      Create a product entity
// @Description  Send the product data to register it in the database.
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        request   body   dto.ProductDTO  true  "Product Data"
// @Success      201 {object} entity.Product
// @Failure      422
// @Failure      400
// @Failure      500
// @Router       /products [post]
// @Security BearerAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var createProductDTO dto.ProductDTO

	err := json.NewDecoder(r.Body).Decode(&createProductDTO)

	if err != nil {
		http.Error(w, "Invalid entry", http.StatusUnprocessableEntity)
		return
	}

	createProductUseCase := h.factory.CreateProductUseCase()
	product, err := createProductUseCase.Execute(createProductDTO)

	if err != nil {
		if httpErr, ok := err.(*httperr.HttpError); ok {
			http.Error(w, httpErr.Message, httpErr.StatusCode)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprintf("/products/%s", product.ID))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// Update Product godoc
// @Summary      Update the product data
// @Description  Send the product data and its ID to update.
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        request   body   dto.ProductDTO  true  "Product Data"
// @Param        id   path   string true  "Product ID"
// @Success      204
// @Failure      422
// @Failure      400
// @Failure      500
// @Router       /products/{id} [put]
// @Security BearerAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Error: invalid ID", http.StatusBadRequest)
		return
	}

	parsedID, err := entityPkg.ParseID(id)

	if err != nil {
		http.Error(w, "Error: invalid ID", http.StatusBadRequest)
		return
	}

	var updateProductDTO dto.ProductDTO

	err = json.NewDecoder(r.Body).Decode(&updateProductDTO)

	if err != nil {
		http.Error(w, "Invalid entry", http.StatusUnprocessableEntity)
		return
	}

	updateProductUseCase := h.factory.UpdateProductUseCase()
	err = updateProductUseCase.Execute(updateProductDTO, parsedID.String())

	if err != nil {
		if httpErr, ok := err.(*httperr.HttpError); ok {
			http.Error(w, httpErr.Message, httpErr.StatusCode)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// Delete Product godoc
// @Summary      Delete the product data
// @Description  Send the product ID to delete it.
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id   path   string true  "Product ID"
// @Success      204
// @Failure      400
// @Failure      500
// @Router       /products/{id} [delete]
// @Security BearerAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Error: invalid ID", http.StatusBadRequest)
		return
	}

	parsedID, err := entityPkg.ParseID(id)

	if err != nil {
		http.Error(w, "Error: invalid ID", http.StatusBadRequest)
		return
	}

	deleteProductUseCase := h.factory.DeleteProductUseCase()
	err = deleteProductUseCase.Execute(parsedID.String())

	if err != nil {
		if httpErr, ok := err.(*httperr.HttpError); ok {
			http.Error(w, httpErr.Message, httpErr.StatusCode)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
