package models

import (
	"auth/pkg/database"
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Id          int    `json:"Id"`
	Username    string `json:"Username"`
	Email       string `json:"Email"`
	IsActivated bool   `json:"IsActivated"`
	Password    []byte `json:"-"`
}

func InsertUserToDB(email, username, password string) (User, Error) {
	res, err := database.DB.Query(fmt.Sprintf(`
		INSERT INTO users 
		(email,username,password)
		VALUES('%s', '%s', '%s')`,
		email, username, password))

	if err != nil {
		log.Println("models -> user -> InsertToDb - 1", err)
		return User{}, Error{Status: 500, Message: "Unhandled user query error"}
	}
	defer res.Close()
	return FindUserByEmail(email)
}

func FindUserByEmail(email string) (User, Error) {
	res, err := database.DB.Query(fmt.Sprintf(`
		SELECT * FROM users 
		WHERE email='%s'`,
		email))
	if err != nil {
		log.Println("models -> admin -> FindByEmail")
		return User{}, Error{Status: 500, Message: "Unhandled user query error"}
	}
	defer res.Close()
	return ParseUserFromQuery(res)
}

func FindUserById(id int) (User, Error) {
	res, err := database.DB.Query(fmt.Sprintf(`
		SELECT * FROM users 
		WHERE id=%d`,
		id))
	if err != nil {
		log.Println("models -> admin -> FindByEmail")
		return User{}, Error{Status: 500, Message: "Unhandled user query error"}
	}
	defer res.Close()
	return ParseUserFromQuery(res)
}

func ParseUserFromQuery(res *sql.Rows) (User, Error) {
	var user User
	for res.Next() {
		err := res.Scan(&user.Id, &user.Email, &user.Username, &user.IsActivated, &user.Password)
		if err != nil {
			log.Println("models -> user -> parseFromQuery -> ", err)
			return User{}, Error{Status: 500, Message: "Unhandled user parse Error"}
		}
	}
	if user.Id == 0 || user.Email == "" || string(user.Password) == "" {
		log.Println("models -> user -> parseFromQuery -> ", "USER not found")
		return User{}, Error{Status: 401, Message: "User not found"}
	}
	return user, Error{Status: 200, Message: ""}
}
