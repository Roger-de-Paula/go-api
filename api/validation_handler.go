package api

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-jwt-api/types"
	"net/http"
)

func validateToken(tknStr string) (*types.Claims, error) {
	claims := &types.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}

func HandleProtectedRoute(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			http.Error(w, "No token", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tokenStr := c.Value
	claims, err := validateToken(tokenStr)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s!", claims.Username)))
}
