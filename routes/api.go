package routes

import (
	"activity_tracker_bot/controllers"
	"activity_tracker_bot/initializers"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
	userController := controllers.NewUserController(initializers.DB)
	authController := controllers.NewAuthController(initializers.DB)
	periodController := controllers.NewPeriodController(initializers.DB)
	activityController := controllers.NewActivityController(initializers.DB)

	api := router.Group("/api/v1")

	AuthRoutes(api, authController)
	UserRoutes(api, userController)
	PeriodRoutes(api, periodController)
	ActivityRoutes(api, activityController)
}
