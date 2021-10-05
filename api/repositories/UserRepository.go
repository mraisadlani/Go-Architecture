package repositories

import (
	"go-architecture-mysql/api/entities"
	"go-architecture-mysql/api/payloads"
)

type UserRepository interface {
	View() ([]entities.User, error)
	FindById(uint) (entities.User, error)
	FindByUsername(string) (entities.User, error)
	Create(payloads.CreateRequest) (bool, error)
	Update(payloads.CreateRequest, uint) (bool, error)
	Delete(uint) (bool, error)
}