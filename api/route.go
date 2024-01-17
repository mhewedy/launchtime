package api

import "github.com/gin-gonic/gin"

func Run(addr ...string) error {
	r := gin.Default()

	r.GET("/forward", getOriginalResponse)
	r.GET("/categories", getCategories)
	r.GET("/categories/:id", getCategoryByID)
	r.GET("/items", searchItems)

	return r.Run(addr...)
}
