package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-tel/src"
	"go-tel/src/backbone/helpers/errors_handler"
	"go-tel/src/backbone/service_layer"
	"go-tel/src/products/adapter/repositories"
	"go-tel/src/products/domain"
	"go-tel/src/products/domain/entities"
	"gorm.io/gorm"
	"net/http"
)

func NewHelloCommand(a string) *domain.HelloCommand {
	helloCmd := &domain.HelloCommand{
		Ali: a,
	}
	helloCmd.SetDependencies([]string{
		"uow",
		"request",
	})
	return helloCmd
}

func HelloHandler(c echo.Context, command service_layer.Command, dependencies map[string]any) error {
	uow := dependencies["uow"].(src.UnitOfWork)

	user4 := uow.Transaction(func(tx *gorm.DB) (any, error) {
		b := repositories.NewGormUserRepo(tx)
		id, err := b.ByID(context.Background(), 4)
		data, err := errors_handler.ErrorExistOrSuccess(id, err)
		return data, err
	})
	user3 := uow.Transaction(func(tx *gorm.DB) (any, error) {
		b := repositories.NewGormUserRepo(tx)
		id, err := b.ByID(context.Background(), 3)
		data, err := errors_handler.ErrorExistOrSuccess(id, err)
		return data, err
	})

	a := []entities.User{
		user3.(entities.User),
		user4.(entities.User),
	}

	return c.JSON(http.StatusOK, a)
}
