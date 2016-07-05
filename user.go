package main

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func findUser(id string) User {
	user := User{}
	rows, _ := database.Query("SELECT id, name FROM user where id = ?", id)
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		user.Id = id
		user.Name = name
	}
	return user
}

func createUser(u User) {
	stmt, _ := database.Prepare("INSERT INTO user(name) values(?)")
	stmt.Exec(u.Name)
}

func destroyUser(id string) {
	stmt, _ := database.Prepare("delete from user where id=?")
	stmt.Exec(id)
}

func updateUser(u User) {
	stmt, _ := database.Prepare("update user set name=? where id=?")
	stmt.Exec(u.Name, u.Id)
}
