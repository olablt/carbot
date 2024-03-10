package main

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var confDir = "./apps/05-hello-db/.conf/"

func main() {
	log.Println("Hello, World!")
	setupPath(confDir)

	taskDB, err := openDB(confDir)
	if err != nil {
		log.Fatal("[DEBUG] error opening database", err)
	}
	id, err := taskDB.insert("Name", "Project")
	if err != nil {
		log.Fatal("[DEBUG] error inserting task", err)
	}
	log.Println("ID:", id)
}

// openDB opens a SQLite database and stores that database in our special spot.
func openDB(path string) (*taskDB, error) {
	db, err := sql.Open("sqlite3", filepath.Join(path, "tasks.db"))
	if err != nil {
		return nil, err
	}
	t := taskDB{db, path}
	if !t.tableExists("tasks") {
		err := t.createTable()
		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}

func setupPath(path string) string {
	if err := initDir(path); err != nil {
		log.Fatal(err)
	}
	return path
}
