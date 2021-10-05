package services

import (
	"go-architecture-mysql/api/entities"
	"go-architecture-mysql/api/payloads"
)

type UserService interface {
	ViewUser() ([]entities.User, error)
	FindById(uint) (entities.User, error)
	FindUser(string) (entities.User, error)
	UpdateUser(payloads.CreateRequest, uint) (bool, error)
	DeleteUser(uint) (bool, error)
}