package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"rollic/pkg/utils"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			utils.HTTPErrorHandler(response, "Token is empty!", http.StatusUnauthorized)
			return
		}

		var secretKey = []byte(os.Getenv("SECRET_KEY"))

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return secretKey, nil
		})

		if err != nil {
			utils.HTTPErrorHandler(response, "Not Authorized", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "data", claims)
			handler.ServeHTTP(response, r.WithContext(ctx))
		} else {
			utils.HTTPErrorHandler(response, "Unauthorized", http.StatusUnauthorized)
		}
	}
}
