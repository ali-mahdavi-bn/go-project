package repositories

import (
	"go-tel/src/backbone/adapter"
	"go-tel/src/products/domain/entities"
	"gorm.io/gorm"
)

type UserRepo interface {
	adapter.AbstractRepository[entities.User]
}

type GormUserRepo struct {
	adapter.AbstractRepository[entities.User]
}

func NewGormUserRepo(db *gorm.DB) UserRepo {
	return &GormUserRepo{AbstractRepository: adapter.NewGormRepository[entities.User](db)}
}
