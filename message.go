package main

type Message struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

type Messages []Message

func findMessage(id string) Message {
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

func getMessages() Messages {
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

func createMessage(m Message) {
	stmt, _ := database.Prepare("INSERT INTO message(name, text) values(?, ?)")
	stmt.Exec(m.Name, m.Text)
}

func destroyMessage(id string) {
	stmt, _ := database.Prepare("delete from message where id=?")
	stmt.Exec(id)
}

func updateMessage(m Message) {
	stmt, _ := database.Prepare("update message set name=?, text=? where id=?")
	stmt.Exec(m.Name, m.Text, m.Id)
}
