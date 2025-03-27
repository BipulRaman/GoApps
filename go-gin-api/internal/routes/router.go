package routes

import (
	"go-gin-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	testController := controllers.TestController{}
	router.GET("/test", testController.GetTest)
	router.POST("/test", testController.CreateTest)

	return router
}
