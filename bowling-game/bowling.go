package bowling_game

const (
	MAX_PINS_PER_FRAME, TOTAL_FRAMES int = 10, 10
)

type Bowling struct {
	rolls []int
}

type score struct {
	total int
	index int
}

func (game *Bowling) Roll(pins int) {
	game.rolls = append(game.rolls, pins)
}

func (game *Bowling) Score() int {
	var frame = 0
	var score score = score{total: 0, index:0}
	for frame < TOTAL_FRAMES {
		if game.isSpare(score.index) {
			score.total += MAX_PINS_PER_FRAME + game.spareBonus(score.index)
			score.index += 2
		} else if game.isStrike(score.index) {
			score.total += MAX_PINS_PER_FRAME + game.strikeBonus(score.index)
			score.index += 1
		} else {
			score.total += game.rolls[score.index] + game.rolls[score.index+1]
			score.index += 2
		}
		frame++
	}
	return score.total
}

func (game *Bowling) isSpare(indexRoll int) bool {
	return game.rolls[indexRoll]+game.rolls[indexRoll+1] == MAX_PINS_PER_FRAME
}

func (game *Bowling) spareBonus(indexRoll int) int {
	return game.rolls[indexRoll+2]
}

func (game *Bowling) isStrike(indexRoll int) bool {
	return game.rolls[indexRoll] == MAX_PINS_PER_FRAME
}

func (game *Bowling) strikeBonus(indexRoll int) int {
	return game.rolls[indexRoll+1] + game.rolls[indexRoll+2]
}
