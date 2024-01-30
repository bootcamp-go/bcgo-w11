package repository_test

import (
	"database/sql"
	"task-testing/internal"
	"task-testing/internal/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func init() {
	// database config connection
	cfg := mysql.Config{
		User:      "root",
		Passwd:    "",
		Addr:      "localhost:3306",
		Net:       "tcp",
		DBName:    "task_test_db",
		ParseTime: true,
	}

	// register txdb driver
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

func TestTaskMySQL_FindByID(t *testing.T) {
	t.Run("success 01 - task was found", func(t *testing.T) {
		// arrange
		// - db
		db, err := sql.Open("txdb", "test_task_mysql_find_by_id_success_01")
		require.NoError(t, err)
		defer db.Close()  // rollback
		// - rollback
		defer func(db *sql.DB) {
			// clean data from tasks table
			_, err := db.Exec("DELETE FROM `tasks`")
			require.NoError(t, err)
			// reset auto increment
			_, err = db.Exec("ALTER TABLE `tasks` AUTO_INCREMENT = 1")
			require.NoError(t, err)
		}(db)
		// - set-up
		err = func(db *sql.DB) error {
			_, err := db.Exec(
				"INSERT INTO `tasks` (`id`, `name`, `description`, `completed`, `created_at`) VALUES (?, ?, ?, ?, ?)",
				10, "task 01", "task 01 description", false, time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
			)
			return err
		}(db)
		require.NoError(t, err)
		// - repository
		rp := repository.NewTaskMySQL(db)

		// act
		task, err := rp.FindByID(10)

		// assert
		expectedTask := internal.Task{
			ID:          10,
			Name:        "task 01",
			Description: "task 01 description",
			Completed:   false,
			CreatedAt:   time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		}
		require.NoError(t, err)
		require.Equal(t, expectedTask, task)
	})

	t.Run("failure 01 - task was not found", func(t *testing.T) {
		// arrange
		// - db
		db, err := sql.Open("txdb", "test_task_mysql_find_by_id_failure_01")
		require.NoError(t, err)
		defer db.Close()
		// - set-up
		// ...
		// - repository
		rp := repository.NewTaskMySQL(db)

		// act
		task, err := rp.FindByID(1)

		// assert
		expectedTask := internal.Task{}
		expectedErr := internal.ErrTaskNotFound
		require.Equal(t, expectedTask, task)
		require.ErrorIs(t, err, expectedErr)
		require.EqualError(t, err, expectedErr.Error())
	})
}

func TestTaskMySQL_Save(t *testing.T) {
	t.Run("success 01 task was saved", func(t *testing.T) {
		// arrange
		// - database: connection
		db, err := sql.Open("txdb", "test_task_mysql_save_success_01")
		require.NoError(t, err)
		defer db.Close()
		// - database: rollback
		defer func(db *sql.DB) {
			// clean data from tasks table
			_, err := db.Exec("DELETE FROM `tasks`")
			require.NoError(t, err)
			// reset auto increment
			_, err = db.Exec("ALTER TABLE `tasks` AUTO_INCREMENT = 1")
			require.NoError(t, err)
		}(db)

		// - database: set-up
		// ...
		// - repository
		rp := repository.NewTaskMySQL(db)

		// act
		task := internal.Task{
			Name:        "task 01",
			Description: "task 01 description",
			Completed:   false,
			CreatedAt:   time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		}
		err = rp.Save(&task)

		// assert
		expectedId := 1
		require.NoError(t, err)
		require.Equal(t, expectedId, task.ID)
	})

	t.Run("failure 01 task was not saved - duplicate entry", func(t *testing.T) {
		// arrange
		// - database: connection
		db, err := sql.Open("txdb", "test_task_mysql_save_success_01")
		require.NoError(t, err)
		defer db.Close()
		// - database: rollback
		defer func(db *sql.DB) {
			// clean data from tasks table
			_, err := db.Exec("DELETE FROM `tasks`")
			require.NoError(t, err)
			// reset auto increment
			_, err = db.Exec("ALTER TABLE `tasks` AUTO_INCREMENT = 1")
			require.NoError(t, err)
		}(db)
		// - database: set-up
		err = func(db *sql.DB) error {
			// insert a task
			_, err := db.Exec(
				"INSERT INTO `tasks` (`id`, `name`, `description`, `completed`, `created_at`) VALUES (?, ?, ?, ?, ?)",
				1, "task 01", "task 01 description", false, time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
			)
			return err
		}(db)
		require.NoError(t, err)
		// - repository
		rp := repository.NewTaskMySQL(db)

		// act
		task := internal.Task{
			Name:        "task 01",
			Description: "task 01 description",
			Completed:   false,
			CreatedAt:   time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		}
		err = rp.Save(&task)

		// assert
		require.ErrorIs(t, err, internal.ErrTaskAlreadyExists)
		require.EqualError(t, err, internal.ErrTaskAlreadyExists.Error())
	})
}

// func ResetDB(t *testing.T, db *sql.DB) {
// 	// clean data from tasks table
// 	_, err := db.Exec("DELETE FROM `tasks`")
// 	require.NoError(t, err)
// 	// reset auto increment
// 	_, err = db.Exec("ALTER TABLE `tasks` AUTO_INCREMENT = 1")
// 	require.NoError(t, err)
// }