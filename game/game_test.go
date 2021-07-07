package game

import (
	"fmt"
	"testing"

	"github.com/smwest87/shining-force-tdd/models"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	wantGame := models.Game{}
	game := NewGame()
	assert.Equal(t, wantGame, game)
}

func TestNewCharacter(t *testing.T) {
	wantPlayer := models.Player{
		Name: "Shepard",
		X:    0,
		Y:    0,
	}

	testPlayer := NewPlayer("Shepard")

	assert.Equal(t, wantPlayer, testPlayer)
}

func TestMoveCharacter(t *testing.T) {
	wantPlayer := models.Player{
		Name: "TestPlayer",
		X:    1,
		Y:    0,
	}

	testPlayer := models.Player{
		Name: "TestPlayer",
		X:    0,
		Y:    0,
	}

	testLevel := models.Level{
		ID:   1234,
		MaxX: 10,
		MaxY: 10,
		OutOfBounds: []models.Coordinates{
			{
				X: 5,
				Y: 5,
			},
		},
	}

	err := testPlayer.Move(1, 0, testLevel)

	assert.NoError(t, err)
	assert.Equal(t, wantPlayer.X, testPlayer.X)
}

func TestMoveCharacter_LessThanZero(t *testing.T) {
	testPlayer := models.Player{
		Name: "TestPlayer",
		X:    0,
		Y:    0,
	}

	testLevel := models.Level{
		ID:   1234,
		MaxX: 10,
		MaxY: 10,
		OutOfBounds: []models.Coordinates{
			{
				X: 5,
				Y: 5,
			},
		},
	}

	expectedErrorMessage := fmt.Sprintf("player cannot exit the map. X: %d", testPlayer.X+-1)

	err := testPlayer.Move(-1, 0, testLevel)
	assert.Error(t, err)
	assert.Equal(t, expectedErrorMessage, err.Error())
}
