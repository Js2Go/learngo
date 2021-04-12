package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"learngo/crawler/frontend/controller"
	"net/http"
	"strconv"
)

var handler = controller.CreateSearchResultNoViewHandler()

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/search", SearchHandler)
	r.Run(":8888")
}

func SearchHandler(c *gin.Context) {
	q := c.Query("q")
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		size = 10
	}
	from, err := strconv.Atoi(c.DefaultQuery("from", "1"))
	if err != nil {
		from = 0
	} else {
		from = (from - 1) * size
	}
	result, err := handler.GetSearchResult(q, from, size)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, result)
}
