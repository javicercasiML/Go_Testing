package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulateHunt_Ok(t *testing.T) {
	// arrange
	server := createServer()

	// act
	t.Run("should create a prey", func(t *testing.T) {
		// arrange
		type request struct {
			Speed float64 `json:"speed"`
		}
		type response struct {
			Success bool    `json:"success"`
			Data    request `json:"data"`
		}

		req, res := createRequestTest(http.MethodPut, "/v1/prey", `{"speed":100.1}`)

		//expectedResponse := response{
		//	Success: true,
		//	Data: request{
		//		Speed: 100.1,
		//	},
		//}
		expectedPrey := request{
			Speed: 100.1,
		}

		// act
		server.ServeHTTP(res, req)
		var r response
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedPrey.Speed, r.Data.Speed)

	})

	t.Run("should create a shark", func(t *testing.T) {
		// arrange
		type request struct {
			XPosition float64 `json:"x_position"`
			YPosition float64 `json:"y_position"`
			Speed     float64 `json:"speed"`
		}
		type response struct {
			Success bool    `json:"success"`
			Data    request `json:"data"`
		}

		req, res := createRequestTest(http.MethodPut, "/v1/shark", `{"x_position":1.2, "y_position" : 200.2, "speed" : 120.1}`)

		expectedResponse := response{
			Success: true,
			Data: request{
				XPosition: 1.2,
				YPosition: 200.2,
				Speed:     120.1,
			},
		}

		// act
		server.ServeHTTP(res, req)
		var r response
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedResponse, r)
	})

	t.Run("should simulate a hunt", func(t *testing.T) {
		// arrange
		type response struct {
			Success bool    `json:"success"`
			Message string  `json:"message"`
			Time    float64 `json:"time"`
		}

		req, res := createRequestTest(http.MethodPost, "/v1/simulate", "")

		expectedResponse := response{
			Success: true,
			Message: "catched prey",
			Time:    10,
		}

		// act
		server.ServeHTTP(res, req)
		var r response
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedResponse, r)
	})

}

func TestSimulateHunt_ErrSpeed(t *testing.T) {
	// arrange
	server := createServer()

	// act
	t.Run("should create a prey", func(t *testing.T) {
		// arrange
		type request struct {
			Speed float64 `json:"speed"`
		}
		type response struct {
			Success bool    `json:"success"`
			Data    request `json:"data"`
		}

		req, res := createRequestTest(http.MethodPut, "/v1/prey", `{"speed":100.1}`)

		//expectedResponse := response{
		//	Success: true,
		//	Data: request{
		//		Speed: 100.1,
		//	},
		//}
		expectedPrey := request{
			Speed: 100.1,
		}

		// act
		server.ServeHTTP(res, req)
		var r response
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedPrey.Speed, r.Data.Speed)

	})

	t.Run("should create a shark", func(t *testing.T) {
		// arrange
		type request struct {
			XPosition float64 `json:"x_position"`
			YPosition float64 `json:"y_position"`
			Speed     float64 `json:"speed"`
		}
		type response struct {
			Success bool    `json:"success"`
			Data    request `json:"data"`
		}

		req, res := createRequestTest(http.MethodPut, "/v1/shark", `{"x_position":1.2, "y_position" : 200.2, "speed" : 10.1}`)

		expectedResponse := response{
			Success: true,
			Data: request{
				XPosition: 1.2,
				YPosition: 200.2,
				Speed:     10.1,
			},
		}

		// act
		server.ServeHTTP(res, req)
		var r response
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedResponse, r)
	})

	t.Run("should simulate a hunt", func(t *testing.T) {
		// arrange
		type response struct {
			Success bool    `json:"success"`
			Message string  `json:"message"`
			Time    float64 `json:"time"`
		}

		req, res := createRequestTest(http.MethodPost, "/v1/simulate", "")

		expectedResponse := response{
			Success: false,
			Message: "could not catch it",
			Time:    0,
		}

		// act
		server.ServeHTTP(res, req)
		var r response
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedResponse, r)
	})

}

func TestSimulateHunt_ErrDistance(t *testing.T) {
	// arrange
	server := createServer()

	// act
	t.Run("should create a prey", func(t *testing.T) {
		// arrange
		type request struct {
			Speed float64 `json:"speed"`
		}
		type response struct {
			Success bool    `json:"success"`
			Data    request `json:"data"`
		}

		req, res := createRequestTest(http.MethodPut, "/v1/prey", `{"speed":100.1}`)

		//expectedResponse := response{
		//	Success: true,
		//	Data: request{
		//		Speed: 100.1,
		//	},
		//}
		expectedPrey := request{
			Speed: 100.1,
		}

		// act
		server.ServeHTTP(res, req)
		var r response
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedPrey.Speed, r.Data.Speed)

	})

	t.Run("should create a shark", func(t *testing.T) {
		// arrange
		type request struct {
			XPosition float64 `json:"x_position"`
			YPosition float64 `json:"y_position"`
			Speed     float64 `json:"speed"`
		}
		type response struct {
			Success bool    `json:"success"`
			Data    request `json:"data"`
		}

		req, res := createRequestTest(http.MethodPut, "/v1/shark", `{"x_position":1000.2, "y_position" : 2000.2, "speed" : 120.1}`)

		expectedResponse := response{
			Success: true,
			Data: request{
				XPosition: 1000.2,
				YPosition: 2000.2,
				Speed:     120.1,
			},
		}

		// act
		server.ServeHTTP(res, req)
		var r response
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedResponse, r)
	})

	t.Run("should simulate a hunt", func(t *testing.T) {
		// arrange
		type response struct {
			Success bool    `json:"success"`
			Message string  `json:"message"`
			Time    float64 `json:"time"`
		}

		req, res := createRequestTest(http.MethodPost, "/v1/simulate", "")

		expectedResponse := response{
			Success: false,
			Message: "could not catch it",
			Time:    0,
		}

		// act
		server.ServeHTTP(res, req)
		var r response
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedResponse, r)
	})

}
