package model

import "github.com/golang-jwt/jwt"

type Login struct {
	IdTeacher string `json:"idTeacher"`
	NameUser  string `json:"nameUser"`
	Password  string `json:"password"`
}

type ResponseStatusLogin struct {
	IdTeacher   string `json:"idTeacher"`
	NameTeacher string `json:"nameTeacher"`
	Surname     string `json:"surname"`
	Response    string `json:"response"`
	Token       string `json:"token"`
	RolId       string `json:"rolId"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
