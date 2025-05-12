package routes

import (
	"activity_tracker_bot/controllers"

	"github.com/gin-gonic/gin"
)

func ActivityRoutes(router *gin.RouterGroup, ActivityController *controllers.DBconnActivityController) {
	router.POST("/activity", ActivityController.StoreActivity)
	router.POST("/user-activity", ActivityController.StoreUserActivity)
}
