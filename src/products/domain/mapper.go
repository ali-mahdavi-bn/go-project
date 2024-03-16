package domain

import (
	"go-tel/src/products/adapter/data_models"
	"go-tel/src/products/domain/entities"
)

func MapNewUserModel(user *entities.User) *data_models.User {
	return &data_models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Password:  user.Password,
	}
}
