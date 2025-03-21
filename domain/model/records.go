package model

import (
	"database/sql"
	"mini-game-go/domain"
	"time"
)

var insertRecord = "INSERT INTO records (player_name, score, difficulty) VALUES (?, ?, ?)"
var selectRecord = "SELECT player_name, score, difficulty, created_at FROM records ORDER BY score DESC LIMIT 10;"

type Record struct {
	Name       string
	Score      int
	CreatedAt  time.Time
	Difficulty domain.Difficulty
}

func InsertRecord(db *sql.DB, name string, score, difficulty int) error {
	_, err := db.Exec(insertRecord, name, score, difficulty)
	if err != nil {
		return err
	}
	return nil
}

func GetRecords(db *sql.DB) ([]Record, error) {
	rows, err := db.Query(selectRecord)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	records := make([]Record, 10)
	i := 0
	for rows.Next() {
		var record Record
		err = rows.Scan(&record.Name, &record.Score, &record.Difficulty, &record.CreatedAt)
		if err != nil {
			return nil, err
		}
		records[i] = record
		i++
	}
	if i < 10 {
		records = records[:i]
	}
	return records, nil
}
