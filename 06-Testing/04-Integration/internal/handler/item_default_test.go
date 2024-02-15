package handler_test

import (
	"database/sql"
	"go-testing/integration/internal/handler"
	"go-testing/integration/internal/repository"
	"go-testing/integration/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func init() {
	// set up the database
	cfg := mysql.Config{
		User:      "root",
		Passwd:    "root",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "items_db",
		ParseTime: true,
	}

	// register the txdb driver
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

// TestItemDefault_FindById tests the FindById handler
func TestItemDefault_FindById(t *testing.T) {
	t.Run("case 01 _ success: find an item by its ID", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "items_db")
		require.NoError(t, err)
		defer db.Close()

		// - set up database
		err = func(db *sql.DB) error {
			_, err := db.Exec("INSERT INTO `items` ( `name`, `description`, `price`) VALUES (?, ?, ?)", "item 01", "description 01", 100.00)
			return err
		}(db)
		require.NoError(t, err)

		// - rollback database
		defer func(db *sql.DB) {
			_, err := db.Exec("DELETE FROM `items`")
			require.NoError(t, err)
			_, err = db.Exec("ALTER TABLE `items` AUTO_INCREMENT = 1")
			require.NoError(t, err)
		}(db)

		// - repository
		rp := repository.NewItemMySQL(db)
		// - service
		sv := service.NewItemDefault(rp)
		// - handler
		hd := handler.NewItemDefault(sv)
		// - router
		r := chi.NewRouter()
		r.Get("/items/{id}", hd.FindById())

		// - request
		req := httptest.NewRequest(http.MethodGet, "/items/1", nil)
		// - response
		res := httptest.NewRecorder()

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedBody := `{"message":"item found","data":{"id":1,"name":"item 01","description":"description 01","price":100}}`

		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 02 _ failure: not found", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "items_db")
		require.NoError(t, err)
		defer db.Close()

		// - repository
		rp := repository.NewItemMySQL(db)
		// - service
		sv := service.NewItemDefault(rp)
		// - handler
		hd := handler.NewItemDefault(sv)
		// - router
		r := chi.NewRouter()
		// - set the route
		r.Get("/items/{id}", hd.FindById())

		// - request
		req := httptest.NewRequest(http.MethodGet, "/items/1", nil)
		// - response
		res := httptest.NewRecorder()

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedBody := `{"message":"item not found", "status":"Not Found"}`

		require.Equal(t, http.StatusNotFound, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})

}
