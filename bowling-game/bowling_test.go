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