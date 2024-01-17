package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type User string
type Item string
type Order map[User][]Item
type Date string

var orders = map[Date]Order{}

func getTodayOrder(c *gin.Context) {

	today := Date(time.Now().Format("2006-01-02"))
	o := orders[today]

	if o == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, o)
	}
}

func addToOrder(c *gin.Context) {

	user := User(c.Query("user"))
	item := Item(c.Query("item"))

	today := Date(time.Now().Format("2006-01-02"))
	order := orders[today]

	if order == nil {
		order = map[User][]Item{}
	}

	if len(order[user]) == 0 {
		order[user] = make([]Item, 0)
	}

	order[user] = append(order[user], item)
	orders[today] = order

	fmt.Println(orders)
}
