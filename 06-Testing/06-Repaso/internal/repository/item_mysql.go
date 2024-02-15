package repository

import (
	"database/sql"
	"go-testing/repaso/internal"
)

// ItemMySQL is a struct that represents the methods that a repository
type ItemMySQL struct {
	db *sql.DB
}

// NewItemMysql returns a new ItemMySQL
func NewItemMysql(db *sql.DB) *ItemMySQL {
	return &ItemMySQL{db}
}

// FindAll returns all items in the database
func (m *ItemMySQL) FindAll(limit int) (i []internal.Item, err error) {
	// execute the query
	rows, err := m.db.Query("SELECT i.`id`, i.`name`, i.`description`, i.`price` FROM `items` i")
	if err != nil {
		return nil, err
	}

	// iterate over the rows
	for rows.Next() {
		var item internal.Item
		// scan the row into the item
		err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price)
		if err != nil {
			return nil, err
		}

		// append the item to the slice
		i = append(i, item)
	}

	// check if there was an error while iterating over the rows
	if rows.Err() != nil {
		err = rows.Err()
	}

	return
}