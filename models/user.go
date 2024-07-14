package models

import (
	"errors"
	"github.com/pol-cova/go-url-shortener-api/db"
	"github.com/pol-cova/go-url-shortener-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u *User) Authenticate() error {
	// Authenticate user logic
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var RetrievedPassword string
	err := row.Scan(&u.ID, &RetrievedPassword)
	if err != nil {
		return err
	}
	isAuthenticated := utils.ValidatePasswordHash(u.Password, RetrievedPassword)
	if !isAuthenticated {
		return errors.New("invalid credentials")
	}
	return nil
}

func (u User) Profile() (User, error) {
	query := `SELECT email FROM users WHERE id = ?`
	row := db.DB.QueryRow(query, u.ID)
	err := row.Scan(&u.Email)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

// Delete account
func (u User) Delete() error {
	query := `DELETE FROM users WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.ID)
	return err
}
