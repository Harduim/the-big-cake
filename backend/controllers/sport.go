package controllers

import (
	"github.com/Harduim/the-big-cake/config"
	"github.com/Harduim/the-big-cake/models"
	"github.com/gin-gonic/gin"
)

type Sport struct{}

func (s *Sport) GetAll(c *gin.Context) {
	var sports []models.Sport
	config.Db.Find(&sports)
	c.JSON(200, sports)
}

func (s *Sport) Update(c *gin.Context) {
	var sport models.Sport
	id := c.Params.ByName("id")
	if err := config.Db.First(&sport, "id = ?", id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	err := c.BindJSON(&sport)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	config.Db.Save(&sport)
	c.JSON(200, sport)
}
