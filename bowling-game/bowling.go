package bowling_game

type Bowling struct {
	rolls []int
}

func (game *Bowling) roll(pins int) {
	game.rolls = append(game.rolls, pins)
}
