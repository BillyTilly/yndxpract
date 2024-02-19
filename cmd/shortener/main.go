package main

import (
	"io"
	"math/rand"
	"net/http"

	"github.com/BillyTilly/yndxpract/config"
	"github.com/gin-gonic/gin"
)

var generatedUrls = make(map[string]string)

func generateURLHandler(c *gin.Context) {
	if c.Request.Header.Get("Content-type") != "text/plain; charset=utf-8" {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(c.Request.Body)
	key := generateKey()
	generatedUrls[key] = string(body)

	answer := config.AppConfig.BaseURL + "/" + key

	c.Writer.Header().Set("Content-type", "text/plain")
	c.Writer.WriteHeader(http.StatusCreated)
	c.Writer.Write([]byte(answer))
}

func redirectHandler(c *gin.Context) {
	key := c.Request.URL.Path[len("/"):]
	if key == "" {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	res, ok := generatedUrls[key]
	if ok {
		c.Writer.Header().Set("Location", res)
		c.Writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}
}

func generateKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func main() {
	config.GetConfig()

	r := gin.New()
	r.POST("/", generateURLHandler)
	r.GET("/:key", redirectHandler)

	r.Run(config.AppConfig.Host)
}
