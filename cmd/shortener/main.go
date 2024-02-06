package main

import (
	"io"
	"math/rand"
	"net/http"
)

var host = "localhost:8080"

var a = make(map[string]string)
var generatedUrl string

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		generateUrlHandler(w, r)
		return
	} else {
		redirectHandler(w, r)
	}
}

func generateUrlHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	key := generateKey()
	a[key] = string(body)
	generatedUrl = key

	answer := "http://localhost:8080/" + key

	w.Header().Set("Content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(answer))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/"):]
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, ok := a[key]
	if ok {
		w.Header().Set("Location", res)
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		w.WriteHeader(http.StatusBadRequest)
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
	err := http.ListenAndServe(host, http.HandlerFunc(handle))
	if err != nil {
		panic(err)
	}
}
