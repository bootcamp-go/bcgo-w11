package handler

import (
	"context"
	"fmt"
	"io"
	"movies-testing/internal"
	"movies-testing/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func NewRequest(method, url string, body io.Reader, urlParams map[string]string, urlQuery map[string]string) *http.Request {
	// old request
	req := httptest.NewRequest(method, url, body)
	
	// url params
	// - new request with a new context with key chi.RouteCtxKey and value chiCtx -> "id":"1"
	if urlParams != nil {
		chiCtx := chi.NewRouteContext()
		for k, v := range urlParams {
			chiCtx.URLParams.Add(k, v)
		}
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
	}

	// url query
	// r.URL.RawQuery = "name=task 1"
	if urlQuery != nil {
		query := req.URL.Query()
		for k, v := range urlQuery {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode() // raw string
	}

	return req
}

func TestTaskDefault_GetTask(t *testing.T) {
	t.Run("success 01 - should return a task", func(t *testing.T) {
		// arrange
		// - repository
		db := map[int]internal.Task{
			1: {
				ID:          1,
				Name:        "task 1",
				Description: "description 1",
				Completed:   false,
			},
		}
		rp := repository.NewTaskMap(db)
		// - handler
		hd := NewTaskDefault(rp)
		hdFunc := hd.GetTask()

		// act
		req := NewRequest("GET", "/tasks/1", nil, map[string]string{"id": "1"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"id":1,"name":"task 1","description":"description 1","completed":false}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("failure 01 - task not found", func(t *testing.T) {
		// arrange
		// - repository
		rp := repository.NewTaskMap(nil)
		// - handler
		hd := NewTaskDefault(rp)
		hdFunc := hd.GetTask()

		// act
		req := NewRequest("GET", "/tasks/1", nil, map[string]string{"id": "1"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusNotFound
		expectedBody := fmt.Sprintf(`{"status":"%s","message":"%s"}`, http.StatusText(expectedCode), "task not found")
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("failure 02 - invalid id", func(t *testing.T) {
		// arrange
		// - repository
		// ...
		// - handler
		hd := NewTaskDefault(nil)
		hdFunc := hd.GetTask()

		// act
		req := httptest.NewRequest("GET", "/tasks/invalid", nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := fmt.Sprintf(`{"status":"%s","message":"%s"}`, http.StatusText(expectedCode), "invalid id")
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}