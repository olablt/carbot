package carbot

import (
	"database/sql"
	"log"
	"path/filepath"
)

// openDB opens a SQLite database and stores that database in our special spot.
func OpenCarbotDB(path string) (*AdsDB, error) {
	db, err := sql.Open("sqlite3", filepath.Join(path, "carbot.db"))
	if err != nil {
		return nil, err
	}
	// t := carbot.AdsDB{db, path}
	adsDB := AdsDB{DB: db, TableName: "ads", FilePath: path}
	if !adsDB.TableExists() {
		log.Println("[DEBUG] Table ads does not exist, creating it")
		err := adsDB.CreateTable()
		if err != nil {
			return nil, err
		}
	} else {
		log.Println("[DEBUG] Table ads exists")
	}
	return &adsDB, nil
}
