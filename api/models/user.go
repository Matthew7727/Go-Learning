package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedpassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	
	result, err := stmt.Exec(u.Email, hashedpassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId

	return err
}

func (u *User) ValidateLogin() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID ,&retrivedPassword)

	if err !=  nil {
		return errors.New("ah ah ah ahhh wrong 2")
	}

	valid := utils.CheckPasswordHash(u.Password, retrivedPassword)

	if !valid {
		return errors.New("ah ah ah ahhh wrong")
	}

	return nil
}