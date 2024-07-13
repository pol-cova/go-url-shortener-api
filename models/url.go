package models

import "time"

type UrlModel struct {
	ID        int64
	Url       string    `binding:"required"`
	Key       string    `binding:"required"`
	CreatedAt time.Time `binding:"required"`
	Clicks    int64
	UserID    int64
}

func (u *UrlModel) Save() error {
	return nil
}
