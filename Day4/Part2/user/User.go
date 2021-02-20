package user

import "encoding/hex"

type User struct {
	Email     string
	Password  string
}

var UserDB = []User{}

type UserJWT struct {
	Email     string
	Password  string
}

var UsersJWT = []UserJWT{}

var APISECRET = hex.EncodeToString([]byte("une-cle-de-32-bytes-de-long-là!"))