package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ____________________________________________________________
// application.go
type ConfigApplication struct {
	Host   string
	Port   int
	DbFile string
}

func NewApplication(cfg ConfigApplication) *Application {
	// default config
	defaultCfg := ConfigApplication{
		Host:   "localhost",
		Port:   8080,
		DbFile: "db.sqlite3",
	}
	if cfg.Host != "" {
		defaultCfg.Host = cfg.Host
	}
	if cfg.Port != 0 {
		defaultCfg.Port = cfg.Port
	}
	if cfg.DbFile != "" {
		defaultCfg.DbFile = cfg.DbFile
	}

	return &Application{
		host:   defaultCfg.Host,
		port:   defaultCfg.Port,
		dbFile: defaultCfg.DbFile,
	}
}

type Application struct {
	host   string
	port   int
	dbFile string
}

func (a *Application) Run() (err error) {
	// dependency injection
	// - router
	fmt.Println("run application")
	fmt.Println("host:", a.host)
	fmt.Println("port:", a.port)
	fmt.Println("dbFile:", a.dbFile)
	// use host
	// use port
	// use dbFile
	return
}

func main() {
	// ____________________________________________________________
	// Config Files (one of all data source to dynamically configure application)
	// - .config/.yaml
	// - .config/.json
	// - .config/.toml
	// cfg := json.NewDecoder(f).Decode(&ConfigApplicationJSON{})
	
	// - .config/.env
	// env file (config file)
	f, err := os.Open(".config/env")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	// read env file
	for {
		line, err := bufio.NewReader(f).ReadString('\n')
		if err != nil {
			break
		}

		// split on "="
		data := strings.Split(line, "=")
		key, value := data[0], data[1]
		os.Setenv(key, value)
	}


	// ____________________________________________________________
	// Os Enviroment Variables (one of all data source to dynamically configure application)
	// env
	host := os.Getenv("ENV_HOST")
	port, err := strconv.Atoi(os.Getenv("ENV_PORT"))
	if err != nil {
		fmt.Println(err)
		return
	}
	dbFile := os.Getenv("ENV_PATH_DB_FILE")

	// app
	cfg := ConfigApplication{
		// Host:   "localhost",
		// Port:   8080,
		// DbFile: "db.sqlite3",
		Host:   host,
		Port:   port,
		DbFile: dbFile,
	}
	app := NewApplication(cfg)

	// run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
