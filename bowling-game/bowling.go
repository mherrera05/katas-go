package bowling_game

type Bowling struct {
	rolls []int
}

func (game *Bowling) roll(pins int) {
	game.rolls = append(game.rolls, pins)
}

func (game *Bowling) score() int {
	var score int = 0
	for _, pins := range game.rolls {
		score += pins
	}
	return score
}
