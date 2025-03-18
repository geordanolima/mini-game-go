package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Conn() *sql.DB {
	db, err := sql.Open("sqlite3", "./records.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getFilesMigration(suffix string) ([]string, error) {
	path, err := filepath.Abs(os.Args[0])
	if err != nil {
		return nil, err
	}
	folder := filepath.Join(filepath.Dir(path), "/database/migrates")
	dir, err := os.Open(folder)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	var sqls []string
	for _, fileInfo := range fileInfos {
		if strings.HasSuffix(fileInfo.Name(), suffix) {
			filePath := filepath.Join(folder, "/", fileInfo.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				return nil, err
			}
			sqls = append(sqls, string(data))

		}
	}
	return sqls, nil
}

func executeMigration(db *sql.DB, suffix string) error {
	sqls, err := getFilesMigration(suffix)
	if err != nil {
		return err
	}
	for _, sql := range sqls {
		_, err := db.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateTables(db *sql.DB) {
	err := executeMigration(db, ".up.sql")
	if err != nil {
		log.Fatal(err)
	}
}

func DropTables(db *sql.DB) {
	err := executeMigration(db, ".down.sql")
	if err != nil {
		log.Fatal(err)
	}
}
