package bowling_game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateABowlingType(test *testing.T) {
	game := new(Bowling)

	var typeof interface{} = game
	_, ok := typeof.(*Bowling)

	assert.Equal(test, true, ok)
}

func TestShouldMakeARoll(test *testing.T) {
	game := new(Bowling)

	game.roll(5)

	assert.Equal(test, 1, len(game.rolls))
}

func TestShouldScore0WhenItFailsAllTheThrows(test *testing.T) {
	game := new(Bowling)

	for i := 0; i < 20; i++ {
		game.roll(0)
	}

	assert.Equal(test, 0, game.score())
}
