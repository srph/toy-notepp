package utils

import (
	"github.com/srph/failbook/models"
)

// Let's pretend we've done authentication.
// For now we'll always return the first user.
func AuthenticatedUser() models.User {
	auth := models.User{}
	models.Instance.First(&auth)
	return auth
}