package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-sqlx/domain"
	"go-gin-sqlx/usecase"
)

type UsersController struct {
	uc usecase.UsersUsecase
}

func NewUsersController(uc usecase.UsersUsecase) *UsersController {
	return &UsersController{uc}
}

func (c UsersController) Login(ctx *gin.Context) {
	var loginReq domain.LoginRequest
	err := ctx.ShouldBindJSON(&loginReq)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if err := loginReq.ValidateUser(); err != nil {

		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err,
		})
		return
	}
	usersLogin, err := c.uc.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "OK",
		"message": "Success login users",
		"data":    usersLogin,
	})
}
