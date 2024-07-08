package bowling_game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var game *Bowling

func TestShouldScore0WhenItFailsAllTheThrows(test *testing.T) {
	game = new(Bowling)

	makeRolls(20, 0)

	assert.Equal(test, 0, game.Score())
}

func TestShouldScore20WhenItKnockDown1PerThrow(test *testing.T) {
	game = new(Bowling)

	makeRolls(20, 1)

	assert.Equal(test, 20, game.Score())
}

func TestShouldScore20WhenItMakesASpareAndGetAnExtraBall(test *testing.T) {
	game = new(Bowling)

	game.Roll(5)
	game.Roll(5)
	game.Roll(5)
	makeRolls(17, 0)

	assert.Equal(test, 20, game.Score())
}

func TestShouldScore20WhenItMakesAStrikeAndGetAnExtraBall(test *testing.T) {
	game = new(Bowling)

	game.Roll(10)
	game.Roll(2)
	game.Roll(3)
	makeRolls(17, 0)

	assert.Equal(test, 20, game.Score())
}

func TestShouldScore300whenItIsAPerfectGame(test *testing.T) {
	game = new(Bowling)

	makeRolls(12, 10)

	assert.Equal(test, 300, game.Score())
}

func TestShouldScore155whenAllTheThrowAreSpare(test *testing.T) {
	game = new(Bowling)

	makeRolls(20, 5)
	game.Roll(10)

	assert.Equal(test, 155, game.Score())
}

func makeRolls(times int, pins int) {
	for i := 0; i < times; i++ {
		game.Roll(pins)
	}
}
