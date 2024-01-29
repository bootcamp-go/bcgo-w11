package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// database
	// - config
	cfg := mysql.Config{
		User: 			   "root",
		Passwd: 		   "",
		Addr: 			   "localhost:3306",
		Net: 			   "tcp",
		DBName: 		   "task_db",
		ParseTime: 	   	   true,
	}
	// - open connection
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/task_db?parseTime=true")
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	// - ping
	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return
	}
}