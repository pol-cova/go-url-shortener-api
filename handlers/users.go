package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/pol-cova/go-url-shortener-api/models"
	"github.com/pol-cova/go-url-shortener-api/utils"
	"github.com/pol-cova/go-url-shortener-api/validators"
	"net/http"
)

func Signup(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error(), "message": "invalid request"})
	}

	// Validate information
	err = validators.AuthValidator(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error(), "message": "something went wrong"})
	}

	// Save user
	err = user.Save()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error(), "message": "could not create user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "user created"})
}

func Login(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error(), "message": "invalid request"})
	}

	err = user.Authenticate()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error(), "message": "invalid credentials"})
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error(), "message": "could not generate token"})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token, "message": "login successful"})

}
