package user

import "github.com/gin-gonic/gin"

type Controller interface {
	Login(ctx *gin.Context)
}

// RegisterRoutesV1 register routes for version 1
func RegisterRoutesV1(router *gin.RouterGroup, controller Controller) {
	router.POST("/login", controller.Login)
}

// RegisterRoutesV2 register routes for version 2
func RegisterRoutesV2(router *gin.RouterGroup, controller Controller) {
	router.POST("/login", controller.Login)
}
