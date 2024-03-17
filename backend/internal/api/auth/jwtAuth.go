package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/rekjef/openchess/internal/api"
	"github.com/rekjef/openchess/internal/types"
	"github.com/rekjef/openchess/pkg/utils"
)

func CreateJWT(account *types.Account) (string, error) {
	claims := types.AuthClaims{
		ID:       account.ID,
		Nickname: account.Nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	secret := os.Getenv("JWT_TOKEN")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func WithJWTAuth(h http.HandlerFunc, s types.Storage) http.HandlerFunc {
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

		claims, ok := token.Claims.(*types.AuthClaims)
		if !ok {
			api.SendError(w, http.StatusBadRequest, errors.New("something is wrong with your auth"))
		}

		// panic(reflect.TypeOf(claims["ID"]))
		if account.Nickname != claims.Nickname {
			api.PermissionDenied(w)
			return
		}

		h(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, types.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := os.Getenv("JWT_TOKEN")

		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
}

func ParseToken(tokenString string) (*types.AuthClaims, error) {
	var claims *types.AuthClaims

	secret := os.Getenv("JWT_TOKEN")
	token, err := jwt.ParseWithClaims(tokenString, &types.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*types.AuthClaims); ok && token.Valid {
			return claims, nil
		}
	}

	return claims, err
}

func getUserAuthInfo(tokenString string) (types.UserAuthInfo, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return types.UserAuthInfo{ID: -1}, err
	}

	return types.UserAuthInfo{ID: claims.ID}, nil
}

func WhoAmI(logger *utils.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			api.MethodNotAllowed(w, r)
			return
		}
		tokenString, err := r.Cookie("x-jwt-token")
		if err != nil {
			api.SendError(w, http.StatusBadRequest, err)
		}

		userInfo, err := getUserAuthInfo(tokenString.Value)
		if err != nil {
			api.SendError(w, http.StatusBadRequest, err)
			return
		}

		utils.Encode[types.UserAuthInfo](w, http.StatusOK, userInfo)
	}
}
