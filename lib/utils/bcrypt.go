package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Make it so that bcrypt just panics,
// who gives a shit about error handling
func Bcrypt(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}