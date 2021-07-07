package game

import (
	"github.com/smwest87/shining-force-tdd/models"
)

func NewGame() models.Game {
	game := models.Game{}
	return game
}

func NewPlayer(name string) models.Player {
	player := models.Player{
		Name: name,
		X:    0,
		Y:    0,
	}
	return player
}
