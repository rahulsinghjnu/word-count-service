package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulsinghjnu/word-count-service/service"
)

func InitRoutes(route *gin.Engine) {
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/wc", service.GetWordCount)
}
