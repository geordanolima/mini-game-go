package engine

import (
	"log"
	"mini-game-go/database"
	"mini-game-go/domain/model"
)

func createDatabase() {
	db := database.Conn()
	database.CreateTables(db)
	db.Close()
}

func saveRecord(game *Game) {
	db := database.Conn()
	err := model.InsertRecord(db, game.User.Name, game.Score.Score, int(game.Dificulty))
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func listRecords() []model.Record {
	db := database.Conn()
	records, err := model.GetRecords(db)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	return records
}
