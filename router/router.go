package router

import (
	"video_api/controller"
	"video_api/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	api := r.Group("/api")
	{
		api.POST("/videoClipEffect", controller.HandleEffectRequest)
		api.POST("/contenateAndPost", controller.HandleConcatenationRequeset)

	}
	r.Static("/static", "./static")
	return r
}
