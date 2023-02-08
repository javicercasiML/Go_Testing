package main

import (
	"api/cmd/server/handlers"
	implmovie "api/internal/movies/implementation"

	"github.com/gin-gonic/gin"
)

func main() {
	// env
	// ...

	// instances
	file := "movies.json"
	st := implmovie.NewStorageFile(file)
	rp := implmovie.NewRepositoryLocal(st)
	sv := implmovie.NewServiceLocal(rp)

	// server
	server := gin.Default()
	// -> routes
	routes := server.Group("/api/v1")
	{
		h := handlers.NewControllerMovie(sv)
		group := routes.Group("/movies")
		group.POST("", h.Create())
	}

	if err := server.Run(); err != nil {
		panic(err)
	}
}