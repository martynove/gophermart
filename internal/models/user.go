package models

type User struct {
	ID       int    `json:"-" db:"id"`
	Login    string `json:"login" binding:"required" db:"login"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}
