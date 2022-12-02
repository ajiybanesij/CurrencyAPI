package Business

import (
	"CurrencyAPI/Helpers"
	"CurrencyAPI/Models"
	"CurrencyAPI/Repositories"
	"errors"
	"fmt"
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

	hash, hashError := Helpers.HashPassword(param.Password)
	if hashError != nil {
		fmt.Println("hash error")
		return errors.New("error")
	}

	user.Name = param.Username
	user.Hash = hash

	createError := user.Create()
	if createError != nil {
		return createError
	}
	return nil
}

func LoginUser(param Models.User) (*string, error) {
	var user Repositories.User
	condition := map[string]interface{}{
		"name": param.Username,
	}
	readError := user.Read(condition)
	if readError != nil {
		return nil, errors.New("user is not exists")
	}

	match := Helpers.CheckPasswordHash(param.Password, user.Hash)
	if !match {
		return nil, errors.New("invalid username or password")
	}

	jwt, jwtError := Helpers.GenerateToken(user.Name)
	if jwtError != nil {
		fmt.Println("generate jwt token error", jwtError)
		return nil, errors.New("login error")
	}

	return &jwt, nil
}
