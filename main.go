package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		fmt.Println("index.tmpl.html", time.Now())
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	go func() {
		for {
			fmt.Println("Current Timestamp:", time.Now())
			time.Sleep(time.Second * 60)
		}
	}()

	router.Run(":" + port)
}
