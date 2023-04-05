//go:build integration

package itest

import (
	"encoding/json"
	"hello-go-rest/internal/service"
	"hello-go-rest/internal/util"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/walkerus/go-wiremock"
)

func TestGetProduct(t *testing.T) {

	wiremockURL := util.GetEnvOrDefault("WIREMOCK_URL", "http://localhost:8081")

	wiremockClient := wiremock.NewClient(wiremockURL)
    defer wiremockClient.Reset()

    wiremockClient.StubFor(wiremock.Get(wiremock.URLPathEqualTo("/products/1")).
		WillReturnJSON(
			&service.Product{
				ID: 1,
				Title: "foo",
				Description: "the product",
			},
			map[string]string{"Content-Type": "application/json"},
			200,
		))

	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/product/1", nil)
	assert.NoError(t, err)

	client := &http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	product := &service.Product{}
	err = json.Unmarshal(body, product)
	assert.NoError(t, err)

	t.Log(string(body))

	assert.Equal(t, 1, product.ID)
	assert.Equal(t, "foo", product.Title)
	assert.Equal(t, "the product", product.Description)
}
