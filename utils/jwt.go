package utils

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(data interface{}, key string, t time.Duration) string {

	jsonData, _ := json.Marshal(data)
	payload := string(jsonData)
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(t).Unix(), 0)),
		Subject:   payload,
	}).SignedString([]byte(key))

	return token
}

func ValidateToken(token string, mapper interface{}, key string) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return err
	}

	for key, val := range claims {
		if key == "sub" {
			json.Unmarshal([]byte(val.(string)), mapper)
		}
	}

	return nil
}
