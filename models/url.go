package models

import (
	"database/sql"
	"github.com/pol-cova/go-url-shortener-api/db"
	"time"
)

type UrlModel struct {
	ID        int64
	Url       string `binding:"required"`
	Key       string
	CreatedAt time.Time
	Clicks    int64
	UserID    int64
}

type ShortUrl struct {
	Key   string `binding:"required" json:"key"`
	Short string `binding:"required" json:"short-url"`
}

func (u *UrlModel) Save() error {
	var query string
	var err error
	var result sql.Result

	if u.UserID == 0 {
		// UserID is not present
		query = `
            INSERT INTO urls (url, key, created_at)
            VALUES (?, ?, ?)`
		stmt, err := db.DB.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		result, err = stmt.Exec(u.Url, u.Key, u.CreatedAt)
	} else {
		// UserID is present
		query = `
            INSERT INTO urls (url, key, created_at, user_id)
            VALUES (?, ?, ?, ?)`
		stmt, err := db.DB.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		result, err = stmt.Exec(u.Url, u.Key, u.CreatedAt, u.UserID)
	}

	if err != nil {
		return err
	}
	u.ID, err = result.LastInsertId()
	return err
}

func GetUrl(key string) (string, bool, error) {
	var url string
	var userID int64
	err := db.DB.QueryRow("SELECT url, user_id FROM urls WHERE key = ?", key).Scan(&url, &userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// No result for the given key
			return "", false, nil
		}
		// An error occurred during the query execution
		return "", false, err
	}
	// Check if userID is not null or 0, return true; otherwise, return false
	userExists := userID != 0
	return url, userExists, nil
}

func UpdateClicks(key string) error {
	_, err := db.DB.Exec("UPDATE urls SET clicks = clicks + 1 WHERE key = ?", key)
	return err
}

func GetAllUrlsByUser(userId int64) ([]UrlModel, error) {
	var urls []UrlModel
	rows, err := db.DB.Query("SELECT id, url, key, created_at, clicks, user_id FROM urls WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var u UrlModel
		err = rows.Scan(&u.ID, &u.Url, &u.Key, &u.CreatedAt, &u.Clicks, &u.UserID)
		if err != nil {
			return nil, err
		}
		urls = append(urls, u)
	}
	return urls, nil
}
