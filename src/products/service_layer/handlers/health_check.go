package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-tel/src"
	"go-tel/src/backbone/helpers/errors_handler"
	"go-tel/src/backbone/service_layer"
	"go-tel/src/products/adapter/repositories"
	"go-tel/src/products/domain/entities"
	"gorm.io/gorm"
	"net/http"
)

type HelloCommand struct {
	service_layer.CommandHandles
}

func NewHelloCommand(request echo.Context) *HelloCommand {
	dependencies := []string{
		"uow",
		"request",
	}
	helloCmd := &HelloCommand{}
	helloCmd.SetRequest(request)
	helloCmd.SetDependencies(dependencies)
	return helloCmd
}

func (c *HelloCommand) CommandHandler(dependencies map[string]any) error {
	uow := dependencies["uow"].(src.GormUnitOfWork)
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
	return c.Request.JSON(http.StatusOK, a)
}
