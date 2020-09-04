package router

import (
	"go-project/internal/interfaces/controller"

	"github.com/gin-gonic/gin"
)

// Initilize はルーティング
func Initilize() {
	router := gin.Default()
	controller := controller.NewController()

	router.GET("/", controller.GetUserIndex)
	router.Run(":8080")
}
