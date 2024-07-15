package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/pol-cova/go-url-shortener-api/db"
	"os"
	"time"
)

func GenerateToken(email string, userId int64) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	secret := os.Getenv("SECRET_KEY")
	// Generate token logic
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
		"iat":    time.Now().Unix(),
	})

	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return 0, errors.New("invalid token")
	}
	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}

func LogoutToken(tokenString string) error {
	// Logout token logic
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return errors.New("invalid token")
	}
	if !parsedToken.Valid {
		return errors.New("invalid token")
	}
	// Blacklist token
	query := `INSERT INTO blacklist(token) VALUES(?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tokenString) // Use tokenString instead of parsedToken
	if err != nil {
		return err
	}

	return nil
}

func IsTokenBlacklisted(tokenString string) bool {
	// Check if the token is blacklisted
	query := `SELECT token FROM blacklist WHERE token = ?`
	row := db.DB.QueryRow(query, tokenString)
	var token string
	err := row.Scan(&token)
	if err != nil {
		return false
	}
	return true
}
