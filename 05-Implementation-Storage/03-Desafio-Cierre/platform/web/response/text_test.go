package response_test

import (
	"app/platform/web/response"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for Text function
func TestText(t *testing.T) {
	t.Run("healthcheck", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		code := http.StatusOK
		body := "pong"
		response.Text(rr, code, body)

		// assert
		expectedHeader := http.Header{"Content-Type": []string{"text/plain; charset=utf-8"}}
		expectedCode := http.StatusOK
		expectedBody := "pong"
		require.Equal(t, expectedHeader, rr.Header())
		require.Equal(t, expectedCode, rr.Code)
		require.Equal(t, expectedBody, rr.Body.String())
	})

	t.Run("empty body", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		code := http.StatusOK
		body := ""
		response.Text(rr, code, body)

		// assert
		expectedHeader := http.Header{"Content-Type": []string{"text/plain; charset=utf-8"}}
		expectedCode := http.StatusOK
		expectedBody := ""
		require.Equal(t, expectedHeader, rr.Header())
		require.Equal(t, expectedCode, rr.Code)
		require.Equal(t, expectedBody, rr.Body.String())
	})
}