package app_version

import "github.com/gin-gonic/gin"

type Controller interface {
	GetListVersion(ctx *gin.Context)
}

// RegisterRoutesV1 register routes for version 1
func RegisterRoutesV1(router *gin.RouterGroup, controller Controller) {
	router.GET("/app_version", controller.GetListVersion)
}
