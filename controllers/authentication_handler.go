package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("notflex")
var tokenName = "token"

type Claims struct {
	ID       string `json:id`
	Name     string `json:"name"`
	UserType int    `json:user_type`
	jwt.StandardClaims
}

func generateToken(w http.ResponseWriter, id string, name string, userType int) {
	tokenExpiryTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		ID:       id,
		Name:     name,
		UserType: userType,
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
	})
}

func resetUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:		tokenName,
		Value:		"",
		Expires:	time.Now(),
		Secure:		false,
		HttpOnly:	true,
	})
}

func AdminAuthenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessType := getUserType(r)
		isValidToken := validateUserToken(w, r, accessType)
		if isValidToken && accessType == 1 {
			next.ServeHTTP(w, r)
		} else {
			sendUnauthorizedResponse(w)
		}
	})
}

func MemberAuthenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessType := getUserType(r)
		isValidToken := validateUserToken(w, r, accessType)
		if isValidToken && accessType == 0 {
			next.ServeHTTP(w, r)
		} else {
			sendUnauthorizedResponse(w)
		}
	})
}

func validateUserToken(w http.ResponseWriter, r *http.Request, accessType int) bool {
	isAccessTokenValid, id, email, userType := validateTokenFromCookies(r)
	fmt.Print(id, email, userType, accessType, isAccessTokenValid)

	if isAccessTokenValid {
		isUserValid := userType == accessType
		if isUserValid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(r *http.Request) (bool, string, string, int) {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.UserType
		}

	}
	return false, "", "", -1
}

func getID(r *http.Request) string {
	if coockie, err := r.Cookie(tokenName); err == nil {
		accessToken := coockie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return accessClaims.ID
		}
	}
	return "Kosong"
}

func getUserType(r *http.Request) int {
	if coockie, err := r.Cookie(tokenName); err == nil {
		accessToken := coockie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return accessClaims.UserType
		}
	}
	return -1
}