package response_test

import (
	"app/platform/web/response"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for JSON function
func TestJSON(t *testing.T) {
	t.Run("200 - status ok", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		code := http.StatusOK
		body := struct{Message string}{Message: "ok"}
		response.JSON(rr, code, body)

		// assert
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedCode := http.StatusOK
		expectedBody := `{"Message":"ok"}`
		require.Equal(t, expectedHeader, rr.Header())
		require.Equal(t, expectedCode, rr.Code)
		require.JSONEq(t, expectedBody, rr.Body.String())
	})

	t.Run("400 - status bad request", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		code := http.StatusBadRequest
		body := struct{Message string}{Message: "bad request"}
		response.JSON(rr, code, body)

		// assert
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedCode := http.StatusBadRequest
		expectedBody := `{"Message":"bad request"}`
		require.Equal(t, expectedHeader, rr.Header())
		require.Equal(t, expectedCode, rr.Code)
		require.JSONEq(t, expectedBody, rr.Body.String())
	})

	t.Run("204 - status no content (body nil)", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		code := http.StatusNoContent
		body := any(nil)
		response.JSON(rr, code, body)

		// assert
		expectedHeader := http.Header{}
		expectedCode := http.StatusNoContent
		expectedBody := ""
		require.Equal(t, expectedHeader, rr.Header())
		require.Equal(t, expectedCode, rr.Code)
		require.Equal(t, expectedBody, rr.Body.String())
	})

	t.Run("500 - status internal server error - internal error (not being able to marshal)", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		code := http.StatusOK
		body := make(chan int)
		response.JSON(rr, code, body)

		// assert
		expectedHeader := http.Header{}
		expectedCode := http.StatusInternalServerError
		expectedBody := ""
		require.Equal(t, expectedHeader, rr.Header())
		require.Equal(t, expectedCode, rr.Code)
		require.Equal(t, expectedBody, rr.Body.String())
	})
}