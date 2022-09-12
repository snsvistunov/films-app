package models

type User struct {
	Id       []uint8 `json:"-" db:"id"`
	Login    string  `json:"login" binding:"required,e164"`
	Password string  `json:"password" binding:"required"`
	Age      byte    `json:"age" binding:"required"`
}

type UserSignIn struct {
	Login    string `json:"login" binding:"required,e164"`
	Password string `json:"password" binding:"required"`
}
