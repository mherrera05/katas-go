package bowling_game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var game *Bowling

func TestCreateABowlingType(test *testing.T) {
	game = new(Bowling)

	var typeof interface{} = game
	_, ok := typeof.(*Bowling)

	assert.Equal(test, true, ok)
}

func TestShouldMakeARoll(test *testing.T) {
	game = new(Bowling)

	game.roll(5)

	assert.Equal(test, 1, len(game.rolls))
}

func TestShouldScore0WhenItFailsAllTheThrows(test *testing.T) {
	game = new(Bowling)

	makeRolls(20, 0)

	assert.Equal(test, 0, game.score())
}

func TestShouldScore20whenItKnockDown1PerThrow(test *testing.T) {
	game = new(Bowling)

	makeRolls(20, 1)

	assert.Equal(test, 20, game.score())
}

func makeRolls(times int, pins int) {
	for i := 0; i < times; i++ {
		game.roll(pins)
	}
}
