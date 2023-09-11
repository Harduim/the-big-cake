package controllers

import (
	"github.com/Harduim/the-big-cake/config"
	"github.com/Harduim/the-big-cake/models"
	"github.com/gin-gonic/gin"
)

type SportCategory struct{}

func (sc *SportCategory) GetAll(c *gin.Context) {
	var sportCategories []models.SportCategory
	config.Db.Find(&sportCategories)
	c.JSON(200, sportCategories)
}
