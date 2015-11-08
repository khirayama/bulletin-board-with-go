package message

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

// TODO: make database global constant
var database, err = sql.Open("sqlite3", "./app.db")

func Find(id string) Message {
	message := Message{}
	rows, _ := database.Query("SELECT id, name, text FROM message where id = ?", id)
	for rows.Next() {
		var id int
		var name string
		var text string
		rows.Scan(&id, &name, &text)
		message.Id = id
		message.Name = name
		message.Text = text
	}
	return message
}

func All() Messages {
	var messages Messages
	rows, _ := database.Query("SELECT id, name, text FROM message")

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

func New(name string, text string) {
	stmt, _ := database.Prepare("INSERT INTO message(name, text) values(?, ?)")
	stmt.Exec(name, text)
}

func (m *Message) Save() {
	// FIXME: check to exist item
	if m.Id != 0 {
		m.Update()
	} else {
		stmt, _ := database.Prepare("INSERT INTO message(name, text) values(?, ?)")
		stmt.Exec(m.Name, m.Text)
	}
}

func Delete(id string) {
	stmt, _ := database.Prepare("delete from message where id=?")
	stmt.Exec(id)
}

func (m *Message) Destroy() {
	stmt, _ := database.Prepare("delete from message where id=?")
	stmt.Exec(m.Id)
}

func (m *Message) Update() {
	stmt, _ := database.Prepare("update message set name=?, text=? where id=?")
	stmt.Exec(m.Name, m.Text, m.Id)
}

func init() {
	database.Exec("create table message(id integer primary key autoincrement, name text, text text)")
}
