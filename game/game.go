package game

import (
	"github.com/smwest87/shining-force-tdd/models"
)

func NewGame(mapID int64) models.Game {
	game := models.Game{
		MapID: mapID,
	}
	return game
}
