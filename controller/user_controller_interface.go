package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}
