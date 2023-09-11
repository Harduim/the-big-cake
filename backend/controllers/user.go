package controllers

import (
	"github.com/Harduim/the-big-cake/config"
	"github.com/Harduim/the-big-cake/models"
	"github.com/gin-gonic/gin"
)

type User struct{}

func (u *User) GetAll(c *gin.Context) {
	var users []models.User
	config.Db.Find(&users)
	c.JSON(200, users)
}
