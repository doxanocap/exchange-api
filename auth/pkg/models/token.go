package models

import (
	"auth/pkg/database"
	"database/sql"
	"fmt"
	"log"
)

type Token struct {
	Id           int    `json:"Id"`
	RefreshToken string `json:"RefreshToken"`
}

func InsertTokenToDB(id int, refreshToken string) (Token, Error) {
	res, err := database.DB.Query(fmt.Sprintf(`
		INSERT INTO tokens 
		(id, refreshToken)
		VALUES(%d, '%s')`,
		id, refreshToken))
	if err != nil {
		log.Println("models -> user -> InsertToDb - 1", err)
		return Token{}, Error{Status: 500, Message: err.Error()}
	}
	defer res.Close()
	return Token{id, refreshToken}, Error{Status: 200, Message: ""}
}

func FindTokenById(id int) (Token, Error) {
	res, err := database.DB.Query(fmt.Sprintf(`
		SELECT * FROM tokens 
		WHERE id = %d`,
		id))
	if err != nil {
		log.Println("models -> token -> FindTokenById - 1", err)
		return Token{}, Error{Status: 500, Message: "Unhandled query error"}
	}
	defer res.Close()
	return ParseTokensFromQuery(res)
}

func FindToken(refreshToken string) (Token, Error) {
	res, err := database.DB.Query(fmt.Sprintf(`
		SELECT * FROM tokens 
		WHERE refreshtoken='%s'
	`, refreshToken))

	if err != nil {
		log.Println("models -> token -> FindToken - 1", err)
		return Token{}, Error{Status: 500, Message: "Unhandled query error"}
	}
	return ParseTokensFromQuery(res)
}

func DeleteToken(refreshToken string) Error {
	res, err := database.DB.Query(fmt.Sprintf(`
		DELETE FROM tokens 
		WHERE refreshToken='%s'
	`, refreshToken))
	if err != nil {
		return Error{Status: 500, Message: "Unhandled error delete token"}
	}
	defer res.Close()
	return Error{Status: 200, Message: ""}
}

func UpdateToken(id int, refreshToken string) Error {
	res, err := database.DB.Query(fmt.Sprintf(`
		UPDATE tokens 
		SET refreshtoken='%s'
		WHERE id='%d'`,
		refreshToken, id))
	if err != nil {
		log.Println("models -> token -> UpdateToken -> ", err)
		return Error{Status: 500, Message: "Unhandled query error"}
	}
	defer res.Close()
	return Error{Status: 200, Message: ""}
}

func ParseTokensFromQuery(res *sql.Rows) (Token, Error) {
	var token Token
	for res.Next() {
		err := res.Scan(&token.Id, &token.RefreshToken)
		if err != nil {
			log.Println("models -> tokens -> ParseTokens ->", err)
			return Token{}, Error{Status: 500, Message: "Unhandled parse error"}
		}
	}

	if token.Id == 0 {
		log.Println("models -> token -> FindTokenById -> Failed to parse Tokens")
		return Token{}, Error{Status: 401, Message: "Token not found"}
	}
	return token, Error{Status: 200, Message: ""}
}
