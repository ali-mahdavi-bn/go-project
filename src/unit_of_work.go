package src

import (
	"fmt"
	"go-tel/src/backbone/infrastructr/databases"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	dbFactory *gorm.DB
}

func NewUnitOfWork() UnitOfWork {
	db, err := databases.SetupDatabase()
	if err != nil {
		fmt.Println(err)
	}
	return UnitOfWork{dbFactory: db}
}

func (uow *UnitOfWork) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (uow *UnitOfWork) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (uow *UnitOfWork) Transaction(f func(tx *gorm.DB) (any, error)) any {
	tx := uow.dbFactory.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	value, errF := f(tx)
	if errF != nil {
		return errF
	}

	return value
}
