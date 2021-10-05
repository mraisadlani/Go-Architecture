package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-architecture-mysql/api/exceptions"
	"go-architecture-mysql/api/middlewares"
	"go-architecture-mysql/api/payloads"
	"go-architecture-mysql/api/securities"
	"go-architecture-mysql/api/services"
	"go-architecture-mysql/api/utils"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	userService services.UserService
}

func CreateUserRoutes(r *gin.RouterGroup, userService services.UserService) {
	userHandler := UserController{
		userService: userService,
	}

	r.Use(middlewares.SetupAuthenticationMiddleware())
	r.GET("/ViewAllUser", userHandler.ViewAllUser)
	r.GET("/FindMe", userHandler.FindMe)
	r.GET("/FindByUsername/:username", userHandler.FindByUser)
	r.POST("/UpdateUser/:user_id", userHandler.UpdateUser)
	r.POST("/DeleteUser/:user_id", userHandler.DeleteUser)
}

// @Summary View All User
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Security BearerAuth
// @Success 200 {object} payloads.Respons
// @Failure 500,400,401,404,403,422 {object} exceptions.Error
// @Router /User/ViewAllUser [get]
func (r *UserController) ViewAllUser(c *gin.Context) {
	get, err := r.userService.ViewUser()

	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}

	payloads.HandleSuccess(c, get, "Get Data Successfully", http.StatusOK)
}

// @Summary Find User
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Security BearerAuth
// @Success 200 {object} payloads.Respons
// @Param username path string true "Username"
// @Failure 500,400,401,404,403,422 {object} exceptions.Error
// @Router /User/FindByUsername/{username} [get]
func (r *UserController) FindByUser(c *gin.Context) {
	username := c.Param("username")

	get, err := r.userService.FindUser(username)

	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}

	payloads.HandleSuccess(c, get, "Get Data Successfully", http.StatusOK)
}

// @Summary Find Me
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Security BearerAuth
// @Success 200 {object} payloads.Respons
// @Failure 500,400,401,404,403,422 {object} exceptions.Error
// @Router /User/FindMe [get]
func (r *UserController) FindMe(c *gin.Context) {
	check, err := securities.ExtractAuthToken(c)

	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}

	get, err := r.userService.FindUser(check.Username)

	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}

	payloads.HandleSuccess(c, get, "Get Data Successfully", http.StatusOK)
}

// @Summary Update User
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Security BearerAuth
// @Param user_id path string true "User ID"
// @Param reqBody body payloads.CreateRequest true "Form Request"
// @Success 200 {object} payloads.Respons
// @Failure 500,400,401,404,403,422 {object} exceptions.Error
// @Router /User/UpdateUser/{user_id} [post]
func (r *UserController) UpdateUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("user_id"))

	var updateData payloads.CreateRequest

	if err := c.ShouldBindJSON(&updateData); err != nil {
		exceptions.EntityException(c, err.Error())
		return
	}

	check := utils.ValidationForm(updateData)

	if check != "" {
		exceptions.BadRequestException(c, check)
		return
	}

	findUser, _ := r.userService.FindById(uint(userId))

	if findUser.Username == "" {
		exceptions.NotFoundException(c, errors.New("User not found").Error())
		return
	}

	hash, err := securities.HashPassword(updateData.Password)

	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}

	updateData.Username = strings.ToLower(strings.ReplaceAll(updateData.Username, " ", ""))
	updateData.Password = strings.ReplaceAll(updateData.Password, " ", "")
	updateData.Password = hash

	get, err := r.userService.UpdateUser(updateData, uint(userId))

	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}

	payloads.HandleSuccess(c, get, "Update Data Successfully", http.StatusOK)
}

// @Summary Delete User
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Security BearerAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} payloads.Respons
// @Failure 500,400,401,404,403,422 {object} exceptions.Error
// @Router /User/DeleteUser/{user_id} [post]
func (r *UserController) DeleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("user_id"))

	get, err := r.userService.DeleteUser(uint(userId))

	if err != nil {
		exceptions.AppException(c, err.Error())
		return
	}

	payloads.HandleSuccess(c, get, "Delete Data Successfully", http.StatusOK)
}
