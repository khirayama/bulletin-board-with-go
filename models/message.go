package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Message struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

type Messages []Message

var db, err = sql.Open("sqlite3", "./app.db")

func GetAllMessages() Messages {
	var messages Messages
	rows, _ := db.Query("SELECT id, name, text FROM message")

	for rows.Next() {
		var id int
		var name string
		var text string
		rows.Scan(&id, &name, &text)
		message := Message{}
		message.Id = id
		message.Name = name
		message.Text = text
		messages = append(messages, message)
	}
	return messages
}

func CreateMessage(message Message) {
	stmt, _ := db.Prepare("INSERT INTO message(name, text) values(?, ?)")
	stmt.Exec(message.Name, message.Text)
}

func DeleteMessage(id string) {
	stmt, _ := db.Prepare("delete from message where id=?")
	stmt.Exec(id)
}

func init() {
	db.Exec("create table message(id integer primary key autoincrement, name text, text text)")
}
