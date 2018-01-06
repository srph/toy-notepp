// 1. Check if auth on every request
// 2. Login
// 3. Logout
// 4. Get User
package auth

import (
	"fmt"
	"errors"
	"github.com/srph/failbook/models"
	"github.com/go-macaron/session"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/macaron.v1"
	// "github.com/volatiletech/authboss"
	// "github.com/volatiletech/authboss/auth"
)

var (
	User *models.User
	SessionStore session.Store
	SessionKey = "_APP_USER_"
)

func Macaron() macaron.Handler {
	return func(ctx *macaron.Context, sess session.Store) {
		SessionStore = sess

		id := sess.Get(SessionKey)

		if id != nil {
			User = fetch("id", id)
		}

		ctx.Next()
	}
}

func Authenticated() bool {
	return User != nil
}

func Guest() bool {
	return User == nil
}

func Logout() {
	_ = SessionStore.Delete(SessionKey)
	User = nil
}

func Login(username, password string) error {
	user := fetch("username", username)
	if user == nil {
		return errors.New("User was not found.")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("Username-password combination was incorrect.")
	}

	SessionStore.Set(SessionKey, user.ID)
	User = user
	return nil
}

func fetch(key string, value interface{}) *models.User {
	user := models.User{}
	where := fmt.Sprintf("%s = ?", key)
	models.Instance.Where(where, value).First(&user)
	return &user
}