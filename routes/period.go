package routes

import (
	"activity_tracker_bot/controllers"

	"github.com/gin-gonic/gin"
)

func PeriodRoutes(router *gin.RouterGroup, periodController *controllers.DbConnPeriodController) {
	router.POST("/period", periodController.StorePeriod)
}
