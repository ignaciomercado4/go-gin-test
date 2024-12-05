package main

import (
	"fmt"
	"net/http"
	"strconv"

	"math/rand"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"route": "/",
		})
	})

	r.GET("/randomNumber", func(c *gin.Context) {

		randomNumber := rand.Intn(1000) + 1

		c.HTML(http.StatusOK, "randomNumber.tmpl", gin.H{
			"randomNumber": randomNumber,
		})
	})

	r.GET("/numberParam/:num", func(c *gin.Context) {

		num, err := strconv.Atoi(c.Param("num"))
		var isGreaterThanFive bool

		if err != nil {
			fmt.Println("Unable to convert num param to string")
		}

		if num < 5 {
			isGreaterThanFive = true
		} else {
			isGreaterThanFive = false
		}

		c.HTML(http.StatusOK, "numParam.tmpl", gin.H{
			"numParam":          num,
			"isGreaterThanFive": isGreaterThanFive,
		})

	})

	r.Run()
}
