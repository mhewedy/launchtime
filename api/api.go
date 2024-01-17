package api

import (
	"github.com/gin-gonic/gin"
	"launchtime/datasource"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

var categories = datasource.GetCategories()

type ItemResp struct {
	Desc  string         `json:"desc"`
	Image string         `json:"image"`
	Name  string         `json:"name"`
	Price int            `json:"price"`
	Sizes []ItemSizeResp `json:"sizes"`
}

type ItemSizeResp struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func normalize(input string) string {
	// replace similar chars
	input = strings.ReplaceAll(input, "أ", "ا")
	input = strings.ReplaceAll(input, "ي", "ى")
	return input
}

func Run() {
	r := gin.Default()

	r.GET("/original-response", getOriginalResponse)

	r.GET("/categories", getCategories)

	r.GET("/categories/:id", getCategoryByID)

	r.GET("/items", searchItems)

	r.Run(":5000")
}

func searchItems(c *gin.Context) {
	name := c.Query("name")
	var l []ItemResp

	for _, cat := range categories.Data {
		for _, item := range cat.Items.Data {
			if strings.Contains(normalize(item.Name.Ar), normalize(name)) ||
				strings.Contains(normalize(item.Description.Ar), normalize(name)) {
				m := ItemResp{
					Name:  item.Name.Ar,
					Desc:  item.Description.Ar,
					Image: item.Image,
					Price: item.Price,
				}

				ls := make([]ItemSizeResp, len(item.Sizes))
				for is, size := range item.Sizes {
					ms := ItemSizeResp{
						Name:  size.Name.Ar,
						Price: size.Price,
					}

					ls[is] = ms
				}
				m.Sizes = ls

				sort.SliceStable(ls, func(i, j int) bool {
					return ls[j].Price > ls[i].Price
				})

				l = append(l, m)
			}
			sort.SliceStable(l, func(i, j int) bool {
				return l[j].Name > l[i].Name
			})
		}
	}
	c.JSON(http.StatusOK, l)
}

func getCategoryByID(c *gin.Context) {
	catID, _ := strconv.Atoi(c.Param("id"))

	for _, cat := range categories.Data {
		if cat.ID == catID {

			l := make([]ItemResp, len(cat.Items.Data))

			for i, item := range cat.Items.Data {
				m := ItemResp{
					Name:  item.Name.Ar,
					Desc:  item.Description.Ar,
					Image: item.Image,
					Price: item.Price,
				}

				ls := make([]ItemSizeResp, len(item.Sizes))
				for is, size := range item.Sizes {
					ms := ItemSizeResp{
						Name:  size.Name.Ar,
						Price: size.Price,
					}

					ls[is] = ms
				}
				m.Sizes = ls

				sort.SliceStable(ls, func(i, j int) bool {
					return ls[j].Price > ls[i].Price
				})

				l[i] = m
			}

			c.JSON(http.StatusOK, l)
			return
		}
	}

	c.AbortWithError(http.StatusNotFound, nil)
}

func getCategories(c *gin.Context) {

	l := make([]map[string]any, len(categories.Data))

	for i, cat := range categories.Data {

		m := make(map[string]any)
		m["id"] = cat.ID
		m["name"] = cat.Name.Ar

		l[i] = m
	}

	sort.SliceStable(l, func(i, j int) bool {
		return l[j]["id"].(int) > l[i]["id"].(int)
	})

	c.JSON(http.StatusOK, l)
}

func getOriginalResponse(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}
