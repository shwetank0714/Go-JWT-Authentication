package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
)


var jwtToken = []byte("secret_key")

// userName : Password matching
var users = map[string]string{
	
	"John" : "wwe16timeChampion",
	"Triple H":"wwe14timeChampion",
	"Edge":"wwe11timeChampion",
	"LA Knight": "megastar",
	"Roman Reigns" : "wwe7timechampion",
}

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
	
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	tokenStr := cookie.Value

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtToken, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute*5)

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	tokenString, err := token.SignedString(jwtToken)

	
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,&http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
	})

	
}

