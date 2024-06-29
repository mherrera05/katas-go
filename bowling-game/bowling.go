package bowling_game

const MAX_PINS_PER_FRAME int = 10

type Bowling struct {
	rolls []int
}

func (game *Bowling) roll(pins int) {
	game.rolls = append(game.rolls, pins)
}

func (game *Bowling) score() int {
	var score int = 0
	for frame := 0; frame < 10; {
		score, frame = game.calculateScore(score, frame)
	}
	return score
}

func (game *Bowling) calculateScore(total int, frame int) (int, int) {
	if game.rolls[frame]+game.rolls[frame+1] == MAX_PINS_PER_FRAME {
		total += MAX_PINS_PER_FRAME + game.rolls[frame+2]
		frame += 2
		return total, frame
	}
	total += game.rolls[frame] + game.rolls[frame+1]
	frame++
	return total, frame
}
