package controllers

import (
	model "activity_tracker_bot/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DbConnPeriodController struct {
	DB *gorm.DB
}

func NewPeriodController(db *gorm.DB) *DbConnPeriodController {
	return &DbConnPeriodController{
		DB: db,
	}
}

func (d *DbConnPeriodController) StorePeriod(c *gin.Context) {

	var body model.Period

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
		return
	}

	period := model.Period{
		Name:         body.Name,
		DaysOfWeek:   body.DaysOfWeek,
		DaysOfMonth:  body.DaysOfMonth,
		MonthsOfYear: body.MonthsOfYear,
	}

	result := d.DB.Create(&period)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": period,
	})
}
