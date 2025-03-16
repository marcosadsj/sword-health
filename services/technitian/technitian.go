package technitian

import "github.com/gin-gonic/gin"

func TechnitianController(httpServer *gin.Engine) {

	routeGroup := httpServer.Group("/technitian")

	routeGroup.POST("/create", Create)
	routeGroup.POST("/read", Read)
	routeGroup.POST("/update", Update)
	routeGroup.POST("/delete", Delete)

}

func Create(ctx *gin.Context) {

}

func Read(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}
