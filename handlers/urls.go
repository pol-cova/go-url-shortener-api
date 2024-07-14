package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/pol-cova/go-url-shortener-api/models"
	"github.com/pol-cova/go-url-shortener-api/utils"
	"net/http"
	"time"
)

func ShortUrl(c echo.Context) error {
	var url models.UrlModel
	err := c.Bind(&url)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	key, err := utils.GenerateKey(utils.KeyLength)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not generate key"})
	}

	url.Key = key
	url.CreatedAt = time.Now().UTC()
	userId := c.Get("userId")
	if userId != nil {
		if id, ok := userId.(int64); ok {
			url.UserID = id
		}
	} else {
		url.UserID = 0
	}

	err = url.Save()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not short the url"})
	}

	shortUrl := models.ShortUrl{
		Key:   url.Key,
		Short: utils.GetShortUrl(url.Key),
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "url shortened", "short-url": shortUrl.Short})
}

func RedirectUrl(c echo.Context) error {
	key := c.Param("key")
	url, isAuth, err := models.GetUrl(key)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "url not found"})
	}
	if isAuth {
		models.UpdateClicks(key)
	}
	return c.Redirect(http.StatusFound, url)
}
