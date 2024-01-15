package handler_test

import (
	"app/10-repaso/internal"
	"app/10-repaso/internal/handler"
	"app/10-repaso/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestItemDefault_FindAll_Handler(t *testing.T) {
	t.Run("success 01 - no query", func(t *testing.T) {
		// arrange
		// - db: map
		db := map[int]internal.Item{
			1: {ID: 1, Name: "item 1", Price: 1.0},
			2: {ID: 2, Name: "item 2", Price: 2.0},
			3: {ID: 3, Name: "item 3", Price: 3.0},
		}
		// - repository: map (in the future a double)
		rp := repository.NewItemMap(db)
		// - handler: default
		hd := handler.NewItemDefault(rp)
		hdFunc := hd.FindAll()

		// act
		req := httptest.NewRequest("GET", "/items", nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"1":{"ID":1,"Name":"item 1","Description":"","Price":1},"2":{"ID":2,"Name":"item 2","Description":"","Price":2},"3":{"ID":3,"Name":"item 3","Description":"","Price":3}}`
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("success 02 - query", func(t *testing.T) {
		// arrange
		// - db: map
		db := map[int]internal.Item{
			1: {ID: 1, Name: "item 1", Price: 1.0},
			2: {ID: 2, Name: "item 2", Price: 2.0},
			3: {ID: 3, Name: "item 3", Price: 2.5},
			4: {ID: 4, Name: "item 4", Price: 3.0},
			5: {ID: 5, Name: "item 5", Price: 4.0},
		}
		// - repository: map (in the future a double)
		rp := repository.NewItemMap(db)
		// - handler: default
		hd := handler.NewItemDefault(rp)
		hdFunc := hd.FindAll()

		// act
		req := httptest.NewRequest("GET", "/items", nil)
		query := req.URL.Query()
		query.Add("price_min", "2.0")
		query.Add("price_max", "3.0")
		req.URL.RawQuery = query.Encode()
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"2":{"ID":2,"Name":"item 2","Description":"","Price":2},"3":{"ID":3,"Name":"item 3","Description":"","Price":2.5},"4":{"ID":4,"Name":"item 4","Description":"","Price":3}}`
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("failure 01 - price min not a float", func(t *testing.T) {
		// arrange
		// - db: map
		// ...
		// - repository: map (in the future a double)
		// ...
		// - handler: default
		hd := handler.NewItemDefault(nil)
		hdFunc := hd.FindAll()

		// act
		req := httptest.NewRequest("GET", "/items", nil)
		query := req.URL.Query()
		query.Add("price_min", "not a float")
		req.URL.RawQuery = query.Encode()
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := `price_min must be a number`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedBody, res.Body.String())
	})

	t.Run("failure 02 - price max not a float", func(t *testing.T) {
		// arrange
		// - db: map
		// ...
		// - repository: map (in the future a double)
		// ...
		// - handler: default
		hd := handler.NewItemDefault(nil)
		hdFunc := hd.FindAll()

		// act
		req := httptest.NewRequest("GET", "/items", nil)
		query := req.URL.Query()
		query.Add("price_max", "not a float")
		req.URL.RawQuery = query.Encode()
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := `price_max must be a number`
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedBody, res.Body.String())
	})

	t.Run("failure 03 - only price min set", func(t *testing.T) {
		// arrange

		// act

		// assert
	
	})

	t.Run("failure 04 - only price max set", func(t *testing.T) {
		// arrange

		// act

		// assert
	
	})
}