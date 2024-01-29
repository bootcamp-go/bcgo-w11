package repository

import (
	"database/sql"
	"errors"
	"task-testing/internal"

	"github.com/go-sql-driver/mysql"
)

// NewTaskMySQL creates a new TaskMySQL
func NewTaskMySQL(db *sql.DB) internal.TaskRepository {
	return &TaskMySQL{db}
}

// TaskMySQL is an implementation of TaskRepository for MySQL
type TaskMySQL struct {
	// db is the database connection
	db *sql.DB
}

// FindByID returns a task by its ID
func (t *TaskMySQL) FindByID(id int) (task internal.Task, err error) {
	// execute the query
	row := t.db.QueryRow("SELECT `id`, `name`, `description`, `completed`, `created_at` FROM `tasks` WHERE `id` = ?", id)
	if row.Err() != nil {
		err = row.Err()
		return
	}

	// serialize the task
	err = row.Scan(&task.ID, &task.Name, &task.Description, &task.Completed, &task.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrTaskNotFound
			return
		}
		return
	}

	return
}

// FindAll returns all tasks
func (t *TaskMySQL) FindAll() (tasks []internal.Task, err error) {
	// execute the query
	rows, err := t.db.Query("SELECT `id`, `name`, `description`, `completed`, `created_at` FROM `tasks`")
	if err != nil {
		return
	}

	// iterate over the rows
	for rows.Next() {
		var task internal.Task
		err = rows.Scan(&task.ID, &task.Name, &task.Description, &task.Completed, &task.CreatedAt)
		if err != nil {
			return
		}
		tasks = append(tasks, task)
	}
	if rows.Err() != nil {
		err = rows.Err()
		return
	}

	return
}

// Save saves a task
func (t *TaskMySQL) Save(task *internal.Task) (err error) {
	// execute the query
	result, err := t.db.Exec(
		"INSERT INTO `tasks` (`name`, `description`, `completed`, `created_at`) VALUES (?, ?, ?, ?)",
		task.Name, task.Description, task.Completed, task.CreatedAt,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrTaskAlreadyExists
				return
			default:
				return
			}
		}
		return
	}

	// get the last inserted ID
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return
	}

	// set the ID of the task
	(*task).ID = int(lastInsertedID)

	return
}

// Update updates a task
func (t *TaskMySQL) Update(task *internal.Task) (err error) {
	// execute the query
	_, err = t.db.Exec(
		"UPDATE `tasks` SET `name` = ?, `description` = ?, `completed` = ? WHERE `id` = ?",
		task.Name, task.Description, task.Completed, task.ID,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrTaskAlreadyExists
				return
			default:
				return
			}
		}
		return
	}

	return
}


// Delete deletes a task
func (t *TaskMySQL) Delete(id int) (err error) {
	// execute the query
	result, err := t.db.Exec("DELETE FROM `tasks` WHERE `id` = ?", id)
	if err != nil {
		return
	}

	// get rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}
	
	// check rows affected
	if rowsAffected != 1 {
		switch {
		case rowsAffected == 0:
			err = internal.ErrTaskNotFound
		default:
			err = errors.New("more than one row affected")
		}
		return
	}

	return
}