package main

import (
	"flag"
	"io"
	"math/rand"
	"net/http"

	"github.com/BillyTilly/yndxpract/config"
	"github.com/gin-gonic/gin"
)

var a = make(map[string]string)
var generatedURL string

func generateURLHandler(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	key := generateKey()
	a[key] = string(body)
	generatedURL = key

	answer := config.AppConfig.Host + "/" + key

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
	var host string
	var resultHost string

	flag.StringVar(&host, "a", "localhost:8080", "host")

	flag.Parse()

	flag.StringVar(&resultHost, "b", "http://"+host, "resulted host")

	flag.Parse()

	config.GenerateConfig(host, resultHost)

	r := gin.New()
	r.POST("/", generateURLHandler)
	r.GET("/:key", redirectHandler)

	r.Run(config.AppConfig.Host)
}
