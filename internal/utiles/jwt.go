package utiles

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Clamis struct {
	Username string `json:"username"`
	UserID   string `json:"userID"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}

func GenerateJWT(username string, userID string, isAdmin bool) (string, error) {
	expirationTime := time.Now().Add(10 * time.Hour)

	claims := &Clamis{
		Username: username,
		UserID:   userID,
		IsAdmin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	// return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))

}
