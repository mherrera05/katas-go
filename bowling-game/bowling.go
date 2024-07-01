package bowling_game

const MAX_PINS_PER_FRAME, TOTAL_FRAMES int = 10, 10

type Bowling struct {
	rolls []int
}

func (game *Bowling) roll(pins int) {
	game.rolls = append(game.rolls, pins)
}

func (game *Bowling) score() int {
	var total, frame, indexRoll int = 0, 0, 0
	for frame < TOTAL_FRAMES {
		if game.isSpare(indexRoll) {
			total += MAX_PINS_PER_FRAME + game.spareBonus(indexRoll)
			indexRoll += 2
		} else if game.isStrike(indexRoll) {
			total += MAX_PINS_PER_FRAME + game.strikeBonus(indexRoll)
			indexRoll += 1
		} else {
			total += game.rolls[indexRoll] + game.rolls[indexRoll+1]
			indexRoll += 2
		}
		frame++
	}
	return total
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
