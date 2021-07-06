package game

import (
	"testing"

	"github.com/smwest87/shining-force-tdd/models"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	wantGame := models.Game{
		MapID: 1234,
	}
	game := NewGame(1234)
	assert.Equal(t, wantGame, game)

}
