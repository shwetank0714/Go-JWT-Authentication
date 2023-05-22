package main

import (
	"net/http"
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
	
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {

}
