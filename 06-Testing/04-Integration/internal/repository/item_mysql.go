package repository

import (
	"database/sql"
	"go-testing/integration/internal"
)

// ItemMySQL is a struct that represents the methods that a repository
type ItemMySQL struct {
	db *sql.DB
}

// NewItemMysql returns a new ItemMySQL
func NewItemMySQL(db *sql.DB) *ItemMySQL {
	return &ItemMySQL{db}
}

// FindAll returns all items in the database
func (m *ItemMySQL) FindById(id int) (i internal.Item, err error) {
	// execute the query
	row := m.db.QueryRow("SELECT `id`, `name`, `description`, `price` FROM `items` WHERE `id` = ?", id)

	// scan the row into the item
	err = row.Scan(&i.ID, &i.Name, &i.Description, &i.Price)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			err = internal.ErrRepositoryNotFound
		default:
			err = internal.ErrRepositoryInternal
		}
		return
	}

	return
}
