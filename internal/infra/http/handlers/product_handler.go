package handlers

import (
	"encoding/json"
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

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
