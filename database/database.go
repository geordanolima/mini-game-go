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

func GetFilesMigration(suffix string, migrationDir string) ([]string, error) {
	dir, err := os.Open(migrationDir)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	sqls := []string{}
	for _, fileInfo := range fileInfos {
		if strings.HasSuffix(fileInfo.Name(), suffix) {
			filePath := filepath.Join(migrationDir, fileInfo.Name()) //Removida a barra inicial
			data, err := os.ReadFile(filePath)
			if err != nil {
				return nil, err
			}
			sqls = append(sqls, string(data))
		}
	}
	return sqls, nil
}

func ExecuteMigration(db *sql.DB, sqls []string) error {
	for _, sql := range sqls {
		_, err := db.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateTables(db *sql.DB) {
	sqls, err := GetFilesMigration(".up.sql", "./database/migrates")
	if err != nil {
		log.Fatal(err)
	}
	err = ExecuteMigration(db, sqls)
	if err != nil {
		log.Fatal(err)
	}
}

func DropTables(db *sql.DB) {
	sqls, err := GetFilesMigration(".down.sql", "./database/migrates")
	if err != nil {
		log.Fatal(err)
	}
	err = ExecuteMigration(db, sqls)
	if err != nil {
		log.Fatal(err)
	}
}
