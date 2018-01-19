// Authee is the authentication package for note++
package authee

import (
	"fmt"
	"errors"
	"github.com/srph/toy-notepp/models"
	"github.com/go-macaron/session"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/macaron.v1"
)

const (
	sessionKey = "_APP_USER_"
)

type Auth struct {
	Session session.Store
	User *models.User
}

// Check if user is authenticated
func (auth *Auth) Check() bool {
	return auth.User != nil
}

// Check if user is a guest
func (auth *Auth) Guest() bool {
	return auth.User == nil
}

// Logout the authenticated user
func (auth *Auth) Logout() {
	_ = auth.Session.Delete(sessionKey)
	auth.User = nil
}

// Authenticate the user with the provided credentials
func (auth *Auth) Login(username, password string) error {
	user, err := fetch("username", username)
	if err != nil {
		return errors.New("User was not found.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("Invalid username-password combination")
	}

	auth.Session.Set(sessionKey, user.ID)
	auth.User = user
	return nil
}

// Macaron middleware
func Macaron() macaron.Handler {
	return func(ctx *macaron.Context, sess session.Store) {
		id := sess.Get(sessionKey)
		auth := Auth{}
		auth.Session = sess

		if id != nil {
			user, err := fetch("id", id)
			if err == nil {
				auth.User = user
			}
		}

		ctx.Map(&auth)

		ctx.Next()
	}
}

// Fetch the user with provided key and value
func fetch(key string, value interface{}) (*models.User, error) {
	user := models.User{}
	where := fmt.Sprintf("%s = ?", key)
	err := models.Instance.Where(where, value).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}