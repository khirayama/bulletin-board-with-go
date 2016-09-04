package main

import (
	"fmt"
	"github.com/markbates/goth"
)

type User struct {
	Id       int    `json:"id"`
	Provider string `json:"provider"`
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
	ImageUrl string `json:"image_url"`
}

func FindOrCreateUserFromAuthHash(gothUser goth.User) User {
	user := User{}

	fmt.Println(gothUser)
	fmt.Println(
		gothUser.Provider,
		gothUser.UserID,
		gothUser.NickName,
		gothUser.AvatarURL,
	)
	rows, _ := database.Query(
		"SELECT id, provider, uid, nickname, image_url FROM users where uid = ? and provider = ?",
		gothUser.UserID,
		gothUser.Provider,
	)
	for rows.Next() {
		var id int
		var provider string
		var uid string
		var nickname string
		var imageUrl string

		rows.Scan(&id, &provider, &uid, &nickname, &imageUrl)

		user.Id = id
		user.Provider = provider
		user.Uid = uid
		user.Nickname = nickname
		user.ImageUrl = imageUrl
		return user
	}

	stmt, _ := database.Prepare("INSERT INTO users(provider, uid, nickname, image_url) values(?, ?, ?, ?)")
	fmt.Println(stmt)
	res, _ := stmt.Exec(
		gothUser.Provider,
		gothUser.UserID,
		gothUser.NickName,
		gothUser.AvatarURL,
	)
	id, _ := res.LastInsertId()
	user.Id = int(id)
	user.Provider = gothUser.Provider
	user.Uid = gothUser.UserID
	user.Nickname = gothUser.NickName
	user.ImageUrl = gothUser.AvatarURL
	return user
}

// func findUser(id string) User {
// 	user := User{}
// 	rows, _ := database.Query("SELECT id, name FROM user where id = ?", id)
// 	for rows.Next() {
// 		var id int
// 		var name string
// 		rows.Scan(&id, &name)
// 		user.Id = id
// 		user.Name = name
// 	}
// 	return user
// }
//
// func createUser(u User) {
// 	stmt, _ := database.Prepare("INSERT INTO user(name) values(?)")
// 	stmt.Exec(u.Name)
// }
//
// func destroyUser(id string) {
// 	stmt, _ := database.Prepare("delete from user where id=?")
// 	stmt.Exec(id)
// }
//
// func updateUser(u User) {
// 	stmt, _ := database.Prepare("update user set name=? where id=?")
// 	stmt.Exec(u.Name, u.Id)
// }
