package handler

import (
	"encoding/json"
	"hello-go-rest/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) GetProduct(writer http.ResponseWriter, request *http.Request) {

	productID, err := strconv.Atoi(chi.URLParam(request, "productId"))
	if err != nil {
		WriteError(writer, err)
		return
	}

	product, err := h.productService.GetProduct(productID)

	if err != nil {
		WriteError(writer, err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(product)

	if err != nil {
		WriteError(writer, err)
		return
	}
}

func WriteError(writer http.ResponseWriter, err error) {
	log.Error(err.Error())
	http.Error(writer, err.Error(), http.StatusInternalServerError)
}
