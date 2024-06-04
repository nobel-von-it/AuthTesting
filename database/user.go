package database

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Pass  string `json:"password"`
}
type Hash interface {
	HashPassword() (string, error)
}

func (u User) HashPassword() (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}
func ComparePassword(hashpass, pass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(pass)); err != nil {
		return false
	}
	return true
}
