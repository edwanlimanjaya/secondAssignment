package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"tugasKedua/model"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("@bebasapasaja_123!")
var tokenName = "token"

type Claims struct {
	IdUser    int    `json : "idUser"`
	Name_user string `json : "name_user"`
	User_type int    `json : "user_type"`
	jwt.StandardClaims
}

func generateToken(w http.ResponseWriter, id int, name string, user_type int) {
	tokenExpiryTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		IdUser:    id,
		Name_user: name,
		User_type: user_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)

	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})
}

func resetUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})
}

func validateTokenFromCookies(r *http.Request) (bool, int, string, int) {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err == nil && parsedToken.Valid {
			return true, accessClaims.IdUser, accessClaims.Name_user, accessClaims.User_type
		}
	}

	return false, -1, "", -1
}

func validateUserToken(r *http.Request, accessType int) bool {
	isAccessTokenValid, id, email, user_type := validateTokenFromCookies(r)
	fmt.Println(id, email, user_type, accessType, isAccessTokenValid)

	if isAccessTokenValid {
		isUserValid := user_type == accessType

		if isUserValid {
			return true
		}
	}

	return false
}

func sendUnAuthorizedResponse(w http.ResponseWriter) {
	var response model.Response
	response.Status = 404
	response.Message = "Unauthorized Access"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Authenticate(next http.HandlerFunc, accessType int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidation := validateUserToken(r, accessType)
		if !isValidation {
			sendUnAuthorizedResponse(w)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
