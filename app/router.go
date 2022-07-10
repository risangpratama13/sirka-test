package app

import (
	"github.com/gin-gonic/gin"
	"github.com/risangpratama13/sirka-test/controller"
)

func UserRouter(router *gin.RouterGroup, controller controller.UserController) {
	router.GET("/DisplayAllUsers", controller.FindAll)
	router.POST("/DisplayUser", controller.FindById)
}
