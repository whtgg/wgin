package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "wgin/docs"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return Router
}
