package game

import (
	"github.com/smwest87/shining-force-tdd/models"
)

func NewGame(characterIDs []int64, mapID int64) models.Game {
	game := models.Game{
		CharacterIDs: characterIDs,
		MapID:        mapID,
		Paused:       false,
	}
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
