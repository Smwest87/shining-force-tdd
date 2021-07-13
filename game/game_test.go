package game

import (
	"fmt"
	"testing"

	"github.com/smwest87/shining-force-tdd/models"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	wantGame := models.Game{
		Paused: false,
		CharacterIDs: []int64{
			1,
		},
		MapID: 1,
	}
	// Need to use a character ID from the character table since Postgres will be generating that ID dynamically
	testCharacter := models.Player{
		ID:   1,
		Name: "Shepard",
		X:    0,
		Y:    0,
	}

	// Need to use a map ID from the map table since Postgres will be generating that ID dynamically
	testLevel := models.Level{
		ID:   1,
		MaxX: 10,
		MaxY: 10,
		OutOfBounds: []models.Coordinates{
			{
				X: 5,
				Y: 5,
			},
			{
				X: 6,
				Y: 5,
			},
		},
	}

	characterIDs := []int64{testCharacter.ID}

	game := NewGame(characterIDs, testLevel.ID)
	assert.Equal(t, wantGame, game)
}

func TestPauseGame(t *testing.T) {
	testGame := models.Game{
		Paused: false,
	}
	err := testGame.PauseGame()
	assert.Nil(t, err)
	assert.Equal(t, true, testGame.Paused)
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

func TestMoveCharacter_XLessThanZero(t *testing.T) {
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

func TestMoveCharacter_YLessThanZero(t *testing.T) {
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

	expectedErrorMessage := fmt.Sprintf("player cannot exit the map. Y: %d", testPlayer.Y+-1)

	err := testPlayer.Move(0, -1, testLevel)
	assert.Error(t, err)
	assert.Equal(t, expectedErrorMessage, err.Error())
}

func TestMoveCharacter_OutOfBounds(t *testing.T) {
	testPlayer := models.Player{
		X: 0,
		Y: 0,
	}
	testLevel := models.Level{
		ID:   1234,
		MaxX: 10,
		MaxY: 10,
		OutOfBounds: []models.Coordinates{
			{
				X: 1,
				Y: 1,
			},
		},
	}

	xMove := 1
	yMove := 1
	expectedErrorMessage := fmt.Sprintf("character cannot enter out of bounds. attempted coordinate: X: %d Y: %d out of bounds coordinate: X: %d Y: %d", testPlayer.X+int64(xMove), testPlayer.Y+int64(yMove), testLevel.OutOfBounds[0].X, testLevel.OutOfBounds[0].Y)

	err := testPlayer.Move(int64(xMove), int64(yMove), testLevel)

	assert.Equal(t, expectedErrorMessage, err.Error())
}

func TestMoveCharacter_PreventDiagonalMovement(t *testing.T) {
	testPlayer := models.Player{
		X: 0,
		Y: 0,
	}
	testLevel := models.Level{
		ID:   1234,
		MaxX: 10,
		MaxY: 10,
		OutOfBounds: []models.Coordinates{
			{
				X: 0,
				Y: 0,
			},
		},
	}
	xMove := 1
	yMove := 1

	expectedErrorMessage := fmt.Sprintf("player cannot move in a diagonal motion: xMove: %d yMove: %d", xMove, yMove)

	err := testPlayer.Move(int64(xMove), int64(yMove), testLevel)

	assert.Error(t, err)
	assert.Equal(t, expectedErrorMessage, err.Error())

}
