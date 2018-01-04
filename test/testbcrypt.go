// Run this to get run get the
// bcrypt-hashed value of `123`.
package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
	fmt.Println(string(hashed))
}