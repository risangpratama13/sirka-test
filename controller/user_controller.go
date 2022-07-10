package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/risangpratama13/sirka-test/controller/request"
	"github.com/risangpratama13/sirka-test/service"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{userService}
}

func (controller *UserControllerImpl) FindById(ctx *gin.Context) {
	var displayUserRequest request.DisplayUserRequest
	err := ctx.Bind(&displayUserRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userid := displayUserRequest.Userid
	user, err := controller.UserService.FindById(ctx.Request.Context(), userid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (controller *UserControllerImpl) FindAll(ctx *gin.Context) {
	users := controller.UserService.FindAll(ctx.Request.Context())
	ctx.JSON(http.StatusOK, users)
}
