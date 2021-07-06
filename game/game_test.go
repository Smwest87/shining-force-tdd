package game

import (
	"testing"

	"github.com/smwest87/shining-force-tdd/models"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	wantGame := models.Game{}
	game := NewGame()
	assert.Equal(t, wantGame, game)

}
