package main

import (
	"github.com/MJRichardson/octopus-random-facts/internal/facts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/fact", func(c *gin.Context) {
		f := facts.RandomFact()
		c.JSON(http.StatusOK, gin.H{
			"title": f.Title,
			"detail": f.Detail,
		})
	})
	r.Run()
}