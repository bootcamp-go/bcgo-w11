package handler_test

import (
	"go-testing/repaso/internal"
	"go-testing/repaso/internal/handler"
	"go-testing/repaso/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestItemDefault_FindAll(t *testing.T) {
	t.Run("success 01 - find some items - completed", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewItemMock()
		rp.FindAllFunc = func() (i []internal.Item, err error) {
			i = []internal.Item{
				{ID: 1, Name: "Item 1", Description: "Description 1", Price: 1.1},
				{ID: 2, Name: "Item 2", Description: "Description 2", Price: 2.2},
				{ID: 3, Name: "Item 3", Description: "Description 3", Price: 3.3},
			}
			return
		}
		// - handler
		hd := handler.NewItemDefault(rp)
		hdFunc := hd.FindAll()

		// act
		// - request
		req := httptest.NewRequest(http.MethodGet, "/items", nil)
		// - response
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `[
			{"id":1,"name":"Item 1","description":"Description 1","price":1.1},
			{"id":2,"name":"Item 2","description":"Description 2","price":2.2},
			{"id":3,"name":"Item 3","description":"Description 3","price":3.3}
		]`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
		// verify registries
		require.Equal(t, 2, rp.MethodCalls["FindAll"])
		require.Equal(t, repository.Args{6}, rp.MethodArgs["FindAll"][0])
		require.Equal(t, repository.Args{7}, rp.MethodArgs["FindAll"][1])
	})

	t.Run("failure 01 - repository error", func(t *testing.T) {
		// arrange
		
		// act
		
		// assert
	})
}