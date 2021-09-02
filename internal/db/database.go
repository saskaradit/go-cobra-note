package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}

	return db.Ping()

}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS studybuddy (
		"idNote" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"word" TEXT,
		"definition" TEXT,
		"category" TEXT
	);`
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	statement.Exec()
	log.Println("StudyBuddy Table created")
}

func InsertNote(word string, definition string, category string) {
	InsertNoteSQL := `INSERT INTO studybuddy (word,definition,category)
	VALUES (?,?,?)`

	statement, err := db.Prepare(InsertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted study note successfully")
}

func DisplatyAllNotes() {
	row, err := db.Query("SELECT * FROM studybuddy ORDER BY word")
	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()

	for row.Next() {
		var idNote int
		var word string
		var definition string
		var category string

		row.Scan(&idNote, &word, &definition, &category)
		log.Println("[", category, "]", word, "-", definition)
	}
}
