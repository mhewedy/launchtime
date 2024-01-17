package api

import "github.com/gin-gonic/gin"

func Run(addr ...string) error {
	r := gin.Default()

	r.GET("/orig", getOriginalResponse)

	r.GET("/categories", getCategories)
	r.GET("/categories/:id", getCategoryByID)
	r.GET("/items", searchItems)

	r.GET("/order", getTodayOrder)

	r.POST("/order", addToOrder)

	return r.Run(addr...)
}
