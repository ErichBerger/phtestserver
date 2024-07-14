package models

import (
	"database/sql"
	"fmt"
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
func (u *UserModel) GetID(username string) (int, error) {
	statement := `select id from User where User.username = ?`

	row := u.DB.QueryRow(statement, username)

	var id int

	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
func (u *UserModel) GetName(id int) (string, error) {
	statement := `select fname, lname from User where User.id = ?`

	row := u.DB.QueryRow(statement, id)

	var fname string
	var lname string
	err := row.Scan(&fname, &lname)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", fname, lname), nil
}

func (*UserModel) ValidateAdmin() (bool, error) {

	return true, nil
}
