package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/internal/database"
	"github.com/rekjef/openchess/internal/types"
)

func CreateJWT(account *types.Account) (string, error) {
	claims := &jwt.MapClaims{
		"ExpiresAt": jwt.NewNumericDate(time.Unix(15000, 0)),
		"nickname":  account.Nickname,
	}

	secret := os.Getenv("JWT_TOKEN")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func WithJWTAuth(h http.HandlerFunc, s database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling JWT auth middleware")
		tokenString := r.Header.Get("x-jwt-token")

		token, err := validateJWT(tokenString)
		if err != nil || !token.Valid {
			api.PermissionDenied(w)
			return
		}

		userID, err := api.GetID(r)
		if err != nil {
			api.PermissionDenied(w)
		}
		account, err := s.GetAccountByID(userID)
		if err != nil {
			api.PermissionDenied(w)
		}

		claims := token.Claims.(jwt.MapClaims)
		// panic(reflect.TypeOf(claims["ID"]))
		fmt.Println(claims["nickname"])
		if account.Nickname != claims["nickname"] {
			api.PermissionDenied(w)
			return
		}

		h(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := os.Getenv("JWT_TOKEN")

		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
}
