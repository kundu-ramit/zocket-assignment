package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kundu-ramit/zocket/controller"
)

func RegisterRoutes(router *gin.Engine) {
	kvController := controller.NewKeyValueController()

	// Auth Routes
	router.POST("/auth", kvController.CreateAuth)

	// Key-Value Routes
	router.POST("/keyvalue", kvController.CreateKeyValue)
	router.GET("/keyvalue/:key", kvController.GetKeyValue)
	router.PUT("/keyvalue/:key", kvController.UpdateKeyValue)
	router.DELETE("/keyvalue/:key", kvController.DeleteKeyValue)
}
