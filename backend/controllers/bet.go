package controllers

import (
	"fmt"

	"github.com/Harduim/the-big-cake/config"
	"github.com/Harduim/the-big-cake/models"
	"github.com/gin-gonic/gin"
)

type Bet struct{}

func (b *Bet) GetAll(c *gin.Context) {
	var bets []models.Bet
	config.Db.Find(&bets)
	c.JSON(200, bets)
}

func (b *Bet) Create(c *gin.Context) {
	fmt.Println("Bet Create")
}

func (b *Bet) Update(c *gin.Context) {
	fmt.Println("Bet Update")
}

func (b *Bet) Delete(c *gin.Context) {
	fmt.Println("Bet Delete")
}
