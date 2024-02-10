package main

import (
	"io"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var host = "localhost:8080"

var a = make(map[string]string)
var generatedURL string

func generateURLHandler(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	key := generateKey()
	a[key] = string(body)
	generatedURL = key

	answer := "http://localhost:8080/" + key

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

	res, ok := a[key]
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
	r := gin.New()
	r.POST("/", generateURLHandler)
	// r.GET("/", redirectHandler)

	r.Run(host)
}
