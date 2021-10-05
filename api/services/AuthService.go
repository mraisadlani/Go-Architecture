package services

import (
	"go-architecture-mysql/api/entities"
	"go-architecture-mysql/api/payloads"
)

type AuthService interface {
	DoLogin(payloads.LoginRequest) (entities.User, error)
	DoRegister(payloads.CreateRequest) (bool, error)
}