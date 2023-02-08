package products

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(db []Product) *gin.Engine {
	// instances
	rp := NewRepository(db)
	sv := NewService(rp)
	hand := NewHandler(sv)

	// server
	server := gin.Default()
	// -> routes
	routes := server.Group("/api/v1/products")

	{
		routes.GET("", hand.GetProducts)
	}

	return server
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

type responseProduct struct {
	ID          string  `json:"id"`
	SellerID    string  `json:"seller_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func TestGetProducts(t *testing.T) {
	// arrange
	db := []Product{
		{ID: "mock", SellerID: "1", Description: "generic product", Price: 123.55},
		{ID: "mock2", SellerID: "2", Description: "generic product2", Price: 1000.1},
	}
	server := createServer(db)

	// act

	t.Run("get products with Unmarshal", func(t *testing.T) {
		// arrange
		req, res := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=1", "")

		expectedProd := []responseProduct{
			{ID: "mock",
				SellerID:    "1",
				Description: "generic product",
				Price:       123.55,
			},
		}

		// act
		server.ServeHTTP(res, req)
		var r []responseProduct
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedProd, r)
	})

	t.Run("get products with Ioutil", func(t *testing.T) {
		// arrange
		req, res := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=1", "")

		expectedResponse := `[{"id":"mock","seller_id":"1","description":"generic product","price":123.55}]`

		// act
		server.ServeHTTP(res, req)
		body, err := io.ReadAll(res.Body)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedResponse, string(body))
	})

	t.Run("product not found", func(t *testing.T) {
		// arrange
		req, res := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=0", "")

		errExpected := `{"error":"seller ID not exist"}`
		// act
		server.ServeHTTP(res, req)
		body, err := io.ReadAll(res.Body)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 500, res.Code)
		assert.Equal(t, errExpected, string(body))
	})

	t.Run("query param not found", func(t *testing.T) {
		// arrange
		req, res := createRequestTest(http.MethodGet, "/api/v1/products", "")

		errExpected := `{"error":"seller_id query param is required"}`
		// act
		server.ServeHTTP(res, req)
		body, err := io.ReadAll(res.Body)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 400, res.Code)
		assert.Equal(t, errExpected, string(body))
	})
}
