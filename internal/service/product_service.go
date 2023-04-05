package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProductService struct {
	baseURL string
}

func NewProductService(baseURL string) *ProductService {
	log.Info("Product BaseURL: ", baseURL)
	return &ProductService{
		baseURL: baseURL,
	}
}

func (s *ProductService) GetProduct(productID int) (*Product, error) {
	url := fmt.Sprintf("%s/products/%d", s.baseURL, productID)
	log.Info("Get URL " + url)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	log.Info("Response Code: ", response.StatusCode)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Could not find product %d", productID)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	product := &Product{}
	err = json.Unmarshal(body, product)

	if err != nil {
		return nil, err
	}

	return product, nil
}
