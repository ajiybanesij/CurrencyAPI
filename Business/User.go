package Business

import (
	"CurrencyAPI/Models"
	"CurrencyAPI/Repositories"
	"errors"
)

func CreateUser(param Models.User) error {
	var user Repositories.User
	condition := map[string]interface{}{
		"name": param.Username,
	}
	readError := user.Read(condition)
	if readError == nil {
		return errors.New("username already exists")
	}
	user.Name = param.Username
	user.Hash = param.Password

	createError := user.Create()
	if createError != nil {
		return createError
	}
	return nil
}
