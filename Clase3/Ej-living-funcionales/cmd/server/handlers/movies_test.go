package handlers

import (
	"api/internal/models"
	"api/internal/movies"
	implmovies "api/internal/movies/implementation"
	"api/internal/movies/mocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ______________________________________________________
// tools
func CreateServerMovie(st movies.Storage) *gin.Engine {
	// instances
	rp := implmovies.NewRepositoryLocal(st)
	sv := implmovies.NewServiceLocal(rp)

	// server
	server := gin.Default()
	// -> routes
	routes := server.Group("/api/v1")
	{
		h := NewControllerMovie(sv)
		group := routes.Group("/movies")
		group.POST("", h.Create())
	}

	return server
}
func NewRequest(method, path, body string) (req *http.Request, res *httptest.ResponseRecorder) {
	// request
	req = httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	// response
	res = httptest.NewRecorder()

	return
}

// ______________________________________________________
// tests
var (
	ErrMock = errors.New("internal error")
)

type responseMovie struct {
	Message string      `json:"message"`
	Data    models.Movie `json:"data"`
}

func TestControllerMovie_Create(t *testing.T) {
	// arrange
	st := mocks.NewStorageMock()
	server := CreateServerMovie(st)

	// act
	t.Run("should create a movie", func(t *testing.T) {
		// arrange
		req, res := NewRequest(http.MethodPost, "/api/v1/movies", `{"title":"A", "rating":8.0, "year":1997}`)

		movie := models.Movie{Title: "A", Rating: 8.0, Year: 1997}

		st.Reset()
		st.On("Read").Return([]models.Movie{}, nil)
		st.On("Write", mock.Anything).Return(nil)

		// act
		server.ServeHTTP(res, req)
		var r responseMovie
		err := json.Unmarshal(res.Body.Bytes(), &r)


		// assert
		assert.NoError(t, err)
		assert.Equal(t, 201, res.Code)
		assert.Equal(t, movie.Year, r.Data.Year)
		assert.True(t, st.AssertExpectations(t))
	})

	t.Run("should return binding error", func(t *testing.T) {
		// arrange
		req, res := NewRequest(http.MethodPost, "/api/v1/movies", ``)

		// act
		server.ServeHTTP(res, req)
		var r Response
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 400, res.Code)
	})

	t.Run("should return required error", func(t *testing.T) {
		// ...
	})

	t.Run("should return invalid error", func(t *testing.T) {
		// arrange
		req, res := NewRequest(http.MethodPost, "/api/v1/movies", `{"title":"A", "rating":100.0, "year":1997}`)

		// act
		server.ServeHTTP(res, req)
		var r Response
		err := json.Unmarshal(res.Body.Bytes(), &r)


		// assert
		assert.NoError(t, err)
		assert.Equal(t, 422, res.Code)
		assert.True(t, st.AssertExpectations(t))
	})

	t.Run("should return not unique error", func(t *testing.T) {
		// arrange
		req, res := NewRequest(http.MethodPost, "/api/v1/movies", `{"title":"A", "rating":8.0, "year":1997}`)

		movie := models.Movie{Title: "A"}

		st.Reset()
		st.On("Read").Return([]models.Movie{movie}, nil)

		// act
		server.ServeHTTP(res, req)
		var r Response
		err := json.Unmarshal(res.Body.Bytes(), &r)


		// assert
		assert.NoError(t, err)
		assert.Equal(t, 409, res.Code)
		assert.True(t, st.AssertExpectations(t))
	})

	t.Run("should return internal error", func(t *testing.T) {
		// arrange
		req, res := NewRequest(http.MethodPost, "/api/v1/movies", `{"title":"A", "rating":8.0, "year":1997}`)

		st.Reset()
		st.On("Read").Return([]models.Movie{}, ErrMock)

		// act
		server.ServeHTTP(res, req)
		var r Response
		err := json.Unmarshal(res.Body.Bytes(), &r)


		// assert
		assert.NoError(t, err)
		assert.Equal(t, 500, res.Code)
		assert.True(t, st.AssertExpectations(t))
	})
}