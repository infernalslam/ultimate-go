package account

import "github.com/infernalslam/simple-api/services"

// CreateAccount must create account in db
func CreateAccount() string {
	return services.CreateAccountInDatabase()
}
