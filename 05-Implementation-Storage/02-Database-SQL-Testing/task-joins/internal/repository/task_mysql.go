package repository

import (
	"database/sql"
	"errors"
	"task-joins/internal"

	"github.com/go-sql-driver/mysql"
)

// TaskMySQL is an implementation of TaskRepository for MySQL
type TaskMySQL struct {
	// db is the database connection
	db *sql.DB
}

// FindByID returns a task by its ID
func (t *TaskMySQL) FindByID(id int) (task internal.Task, err error) {
	// execute the query
	query := "SELECT `id`, `name`, `description`, `completed`, `created_at`, `owner_id` FROM `tasks` WHERE `id` = ?"
	row := t.db.QueryRow(query, id)
	if row.Err() != nil {
		err = row.Err()
		return
	}

	// serialize the task
	err = row.Scan(&task.ID, &task.Name, &task.Description, &task.Completed, &task.CreatedAt, &task.OwnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrTaskNotFound
			return
		}
		return
	}

	return
}

// FindByIDWithOwner returns a task by its ID with its owner
func (t *TaskMySQL) FindByIDWithOwner(id int) (task internal.TaskWithOwner, err error) {
	// execute the query
	query := "SELECT t.`id`, t.`name`, t.`description`, t.`completed`, t.`created_at`, o.`name`, o.`email` FROM `tasks` t " +
		     "INNER JOIN `owners` o ON t.`owner_id` = o.`id` WHERE t.`id` = ?"
	row := t.db.QueryRow(query, id)
	if row.Err() != nil {
		err = row.Err()
		return
	}

	// serialize the task
	err = row.Scan(&task.ID, &task.Name, &task.Description, &task.Completed, &task.CreatedAt, &task.OwnerName, &task.OwnerEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrTaskNotFound
			return
		}
		return
	}

	return
}

// Save saves a task
func (t *TaskMySQL) Save(task *internal.Task) (err error) {
	// execute the query
	query := "INSERT INTO `tasks` (`name`, `description`, `completed`, `created_at`, `owner_id`) VALUES (?, ?, ?, ?, ?)"
	result, err := t.db.Exec(query, task.Name, task.Description, task.Completed, task.CreatedAt, task.OwnerID)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrTaskAlreadyExists
				return
			case 1452:
				err = internal.ErrTaskHasNoOwner
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