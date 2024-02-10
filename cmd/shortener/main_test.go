package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateUrlHandler(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
	}
	tests := []struct {
		name    string
		request string
		body    string
		want    want
	}{
		{
			name: "simple test #1",
			body: "http://yandex.ru",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusCreated,
			},
			request: "/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.request, bytes.NewBufferString(tt.body))
			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)

			c.Request = request

			generateURLHandler(c)

			result := w.Result()

			assert.Equal(t, tt.want.statusCode, c.Writer.Status())
			assert.Equal(t, tt.want.contentType, c.Writer.Header().Get("Content-Type"))

			_, err := io.ReadAll(result.Body)
			require.NoError(t, err)
			err = result.Body.Close()
			require.NoError(t, err)
		})
	}
}

func TestRedirectHandler(t *testing.T) {
	type want struct {
		statusCode int
		headerLoc  string
	}
	tests := []struct {
		name    string
		request string
		want    want
	}{
		{
			name: "simple test #1",
			want: want{
				statusCode: http.StatusTemporaryRedirect,
				headerLoc:  "http://yandex.ru",
			},
			request: "/" + generatedURL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.request, nil)
			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)

			c.Request = request

			redirectHandler(c)

			result := w.Result()

			c.Writer.Header().Get("Location")
			assert.Equal(t, tt.want.statusCode, c.Writer.Status())
			assert.Equal(t, tt.want.headerLoc, c.Writer.Header().Get("Location"))

			_, err := io.ReadAll(result.Body)
			require.NoError(t, err)
			err = result.Body.Close()
			require.NoError(t, err)
		})
	}
}
