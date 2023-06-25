package repository

import (
	"auth/pkg/database"
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    []byte `json:"-"`
}

func InsertUserToDB(email, username, password string) (User, Error) {
	res, err := database.DB.Query(fmt.Sprintf(`
		INSERT INTO users 
		(email,username,password)
		VALUES('%s', '%s', '%s')`,
		email, username, password))

	if err != nil {
		log.Println("model -> user -> InsertToDb - 1", err)
		return User{}, Error{Status: 500, Message: "unhandled query error"}
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
		log.Println("model -> admin -> FindByEmail ->", err.Error())
		return User{}, Error{Status: 500, Message: "aurh user query error"}
	}
	return ParseUserFromQuery(res)
}

func FindUserById(id int) (User, Error) {
	res, err := database.DB.Query(fmt.Sprintf(`
		SELECT * FROM users 
		WHERE id=%d`,
		id))
	if err != nil {
		log.Println("model -> admin -> FindByEmail")
		return User{}, Error{Status: 500, Message: "unhandled user query error"}
	}
	defer res.Close()
	return ParseUserFromQuery(res)
}

func ParseUserFromQuery(res *sql.Rows) (User, Error) {
	defer res.Close()

	var user User
	for res.Next() {
		err := res.Scan(&user.Id, &user.Email, &user.Username, &user.PhoneNumber, &user.Password)
		if err != nil {
			log.Println("model -> user -> parseFromQuery -> ", err)
			return User{}, Error{Status: 500, Message: "unhandled user parse error"}
		}
	}

	if user.Id == 0 || user.Email == "" || string(user.Password) == "" {
		log.Println("model -> user -> parseFromQuery -> ", "USER not found")
		return User{}, Error{Status: 401, Message: "user not found"}
	}
	return user, Error{Status: 200, Message: ""}
}
