package models

import (
	"database/sql"
)

type User struct {
	username string
	password string
}

type UserModel struct {
	DB *sql.DB
}

func (*UserModel) ValidateProvider() (bool, error) {
	//DB.

	return true, nil
}

func (*UserModel) ValidateAdmin() (bool, error) {

	return true, nil
}
