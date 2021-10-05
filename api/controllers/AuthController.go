package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-architecture-mysql/api/exceptions"
	"go-architecture-mysql/api/payloads"
	"go-architecture-mysql/api/securities"
	"go-architecture-mysql/api/services"
	"go-architecture-mysql/api/utils"
	"net/http"
	"strings"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
}

func CreateAuthRoutes(r *gin.RouterGroup, authService services.AuthService, userService services.UserService) {
	authHandler := AuthController{
		authService: authService,
		userService: userService,
	}

	r.POST("/Register", authHandler.DoRegister)
	r.POST("/Login", authHandler.DoLogin)
}

// @Summary Register User
// @Description REST API User
// @Accept json
// @Produce json
// @Tags Auth Controller
// @Param reqBody body payloads.CreateRequest true "Form Request"
// @Success 200 {object} payloads.Respons
// @Failure 500,400,401,404,403,422 {object} exceptions.Error
// @Router /Auth/Register [post]
func (r *AuthController) DoRegister(c *gin.Context) {
	var register payloads.CreateRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		exceptions.EntityException(c, err.Error())
		return
	}

	check := utils.ValidationForm(register)

	if check != "" {
		exceptions.BadRequestException(c, check)
		return
	}

	findUser, _ := r.userService.FindUser(register.Username)

	if findUser.Username != "" {
		exceptions.NotFoundException(c, errors.New("Username already exists").Error())
		return
	}

	hash, err := securities.HashPassword(register.Password)

	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}

	register.Username = strings.ToLower(strings.ReplaceAll(register.Username, " ", ""))
	register.Password = strings.ReplaceAll(register.Password, " ", "")
	register.Password = hash

	get, err := r.authService.DoRegister(register)

	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}

	payloads.HandleSuccess(c, get, "Register Successfully", http.StatusOK)
}

// @Summary Login User
// @Description REST API User
// @Accept json
// @Produce json
// @Tags Auth Controller
// @Param reqBody body payloads.LoginRequest true "Form Request"
// @Success 200 {object} payloads.Respons
// @Failure 500,400,401,404,403,422 {object} exceptions.Error
// @Router /Auth/Login [post]
func (r *AuthController) DoLogin(c *gin.Context) {
	var login payloads.LoginRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		exceptions.EntityException(c, err.Error())
		return
	}

	check := utils.ValidationForm(login)

	if check != "" {
		exceptions.BadRequestException(c, check)
		return
	}

	findUser, _ := r.userService.FindUser(login.Username)

	if findUser.Username != "" {
		hashPwd := findUser.Password
		pwd := login.Password

		hash := securities.VerifyPassword(hashPwd, pwd)

		if hash == nil {
			token, err := securities.GenerateToken(findUser.Username)

			if err != nil {
				exceptions.AppException(c, err.Error())
				return
			}

			payloads.ResponseToken(c, "Login Successfully", token, http.StatusOK)
		} else {
			exceptions.BadRequestException(c, errors.New("Password dont matched").Error())
			return
		}
	} else {
		exceptions.NotFoundException(c, errors.New("Username not found").Error())
		return
	}
}