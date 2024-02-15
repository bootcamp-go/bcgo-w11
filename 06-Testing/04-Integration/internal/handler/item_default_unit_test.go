package handler_test

import (
	"context"
	"go-testing/integration/internal"
	"go-testing/integration/internal/handler"
	"go-testing/integration/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

// TestItemDefault_FindById tests the FindById handler
func TestItemDefaultUnit_FindByID(t *testing.T) {
	t.Run("case 01 _ success: find an item by its ID", func(t *testing.T) {
		// arrange
		// - service
		sv := service.NewItemDefaultMock()
		// - handler
		hd := handler.NewItemDefault(sv)
		// handlerFunc := hd.FindById()

		// - set the findbyid mock
		sv.On("FindById", 1).Return(internal.Item{
			ID:          1,
			Name:        "item 01",
			Description: "description 01",
			Price:       100.00,
		}, nil)

		// - request
		req := httptest.NewRequest(http.MethodGet, "/items/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// act
		hd.FindById()(res, req)
		// handlerFunc(res, req)

		// assert
		expectedBody := `{"message":"item found","data":{"id":1,"name":"item 01","description":"description 01","price":100}}`

		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}
