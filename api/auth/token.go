package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userid uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = userid
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {

}
