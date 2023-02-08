package handlers

import (
	"api/internal/movies"
	"errors"

	"github.com/gin-gonic/gin"
)

// constructor: create a new controller
func NewControllerMovie(sv movies.Service) *ControllerMovie {
	return &ControllerMovie{sv: sv}
}

// controller
type ControllerMovie struct {
	// service
	sv movies.Service
}

func (c *ControllerMovie) Create() gin.HandlerFunc {
	type request struct {
		Title  string  `json:"title"`
		Rating float64 `json:"rating"`
		Year   int     `json:"year"`
	}
	
	return func(ctx *gin.Context) {
		// request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ResponseErr(ctx, 400, "could not bind request")
			return
		}
		
		// process
		mv, err := c.sv.Create(req.Title, req.Rating, req.Year)
		if err != nil {
			if errors.Is(err, movies.ErrServiceInvalid) {
				ResponseErr(ctx, 422, "movie not valid")
				return
			}
			if errors.Is(err, movies.ErrServiceNotUnique) {
				ResponseErr(ctx, 409, "movie not unique")
				return
			}
			ResponseErr(ctx, 500, "internal error")
			return
		}


		// response
		ResponseOk(ctx, 201, "movie created", mv)
	}
}