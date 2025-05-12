package controllers

import (
	model "activity_tracker_bot/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DBconnActivityController struct {
	DB *gorm.DB
}

func NewActivityController(db *gorm.DB) *DBconnActivityController {
	return &DBconnActivityController{
		DB: db,
	}
}

func (a *DBconnActivityController) StoreActivity(c *gin.Context) {

	var body model.Activity

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	activity := model.Activity{
		Name:     body.Name,
		PeriodID: body.PeriodID,
	}

	result := a.DB.Create(&activity)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": activity,
	})
}

func (a *DBconnActivityController) StoreUserActivity(c *gin.Context) {

	var body model.UserActivity

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user_activity := model.UserActivity{
		ActivityID: body.ActivityID,
		UserID:     body.UserID,
	}

	result := a.DB.Create(&user_activity)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": user_activity,
	})
}
