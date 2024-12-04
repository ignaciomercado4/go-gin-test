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

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "I'm alive",
		})
	})

	r.GET("/randomNumber", func(c *gin.Context) {

		randomNumber := rand.Intn(1000) + 1

		c.JSON(http.StatusOK, gin.H{
			"randomNumber": randomNumber,
		})
	})

	r.GET("/numberParam/:num", func(c *gin.Context) {

		num, err := strconv.Atoi(c.Param("num"))
		var isGreaterThanFive bool

		if err != nil {
			fmt.Println("Unable to convert num param to string")
		}

		if num > 5 {
			isGreaterThanFive = true
		} else {
			isGreaterThanFive = false
		}

		c.JSON(http.StatusOK, gin.H{
			"numParam":          num,
			"isGreaterThanFive": isGreaterThanFive,
		})

	})

	r.Run()
}
