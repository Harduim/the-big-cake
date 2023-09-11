package controllers

import (
	"github.com/Harduim/the-big-cake/config"
	"github.com/Harduim/the-big-cake/models"
	"github.com/gin-gonic/gin"
)

type Country struct{}

func (ctr *Country) GetAll(c *gin.Context) {
	var countries []models.Country
	config.Db.Find(&countries)
	c.JSON(200, countries)
}
