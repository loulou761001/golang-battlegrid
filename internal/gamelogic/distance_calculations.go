package gamelogic

import (
	"battle-sim/internal/state"
	"battle-sim/internal/types"
	"math"
)

func manhattanDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1)))
}

func FindClosestUnit(unit types.Unit) (types.Unit, int) {
	var closest types.Unit
	minDist := math.MaxInt

	for _, other := range state.Units {
		if IsUnitAlly(&unit, &other) { // éviter de se comparer à soi-même et avec les morts
			continue
		}
		dist := manhattanDistance(unit.X, unit.Y, other.X, other.Y)
		if dist < minDist {
			minDist = dist
			closest = other
		}
	}

	return closest, minDist
}
