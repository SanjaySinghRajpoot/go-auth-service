package routes

import (
	"github.com/SanjaySinghRajpoot/backend/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	// All user routes
	router.GET("/user", controller.GetUsers)
	router.POST("/user/login", controller.LoginUser)
	router.GET("/user/logout", controller.LogoutUser)
	router.POST("/user/refresh", controller.RefreshUser)
}
