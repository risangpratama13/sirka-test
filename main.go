package main

import (
	"github.com/gin-gonic/gin"
	"github.com/risangpratama13/sirka-test/app"
	"github.com/risangpratama13/sirka-test/controller"
	"github.com/risangpratama13/sirka-test/repository"
	"github.com/risangpratama13/sirka-test/service"
	"github.com/risangpratama13/sirka-test/util"
)

func main() {
	r := gin.Default()
	config, err := util.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db := app.NewDB(config)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserControllerImpl(userService)

	routerGroup := r.Group("/MyWeb")
	app.UserRouter(routerGroup, userController)
	r.Run()
}
