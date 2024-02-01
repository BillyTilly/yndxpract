package main

import (
	"io"
	"math/rand"
	"net/http"
)

var host = "localhost:8080"

var a = make(map[string]string)

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		body, _ := io.ReadAll(r.Body)
		key := generateKey()
		a[key] = string(body)

		answer := "http://localhost:8080/" + key

		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(answer))

		return
	} else {
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
