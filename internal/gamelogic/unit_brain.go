package gamelogic

import (
	"battle-sim/internal/state"
	"battle-sim/internal/types"
	"math"
	"math/rand"
)

func moveUnitTowardsClosest(unit *types.Unit) bool {
	closest, dist := FindClosestUnit(*unit)
	if dist <= unit.UnitType.Stats.Speed {
		return false // pas de cible ou déjà sur place
	}

	speed := unit.UnitType.Stats.Speed
	dx := closest.X - unit.X
	dy := closest.Y - unit.Y

	// Normalisation du mouvement (on avance d'au plus "speed" cases)
	moveX := int(math.Copysign(math.Min(math.Abs(float64(dx)), float64(speed)), float64(dx)))
	moveY := int(math.Copysign(math.Min(math.Abs(float64(dy)), float64(speed)), float64(dy)))

	// ⚖ Si speed permet de bouger plus que la distance dans un axe, redistribuer vers l’autre
	remainingSpeed := speed - int(math.Max(math.Abs(float64(moveX)), math.Abs(float64(moveY))))
	if remainingSpeed > 0 {
		if math.Abs(float64(dx)) > math.Abs(float64(dy)) {
			extraY := int(math.Copysign(math.Min(float64(remainingSpeed), math.Abs(float64(dy)-math.Abs(float64(moveY)))), float64(dy)))
			moveY += extraY
		} else {
			extraX := int(math.Copysign(math.Min(float64(remainingSpeed), math.Abs(float64(dx)-math.Abs(float64(moveX)))), float64(dx)))
			moveX += extraX
		}
	}
	newX := unit.X + moveX
	newY := unit.Y + moveY

	for _, other := range state.Units {
		if other.Id != unit.Id && other.X == newX && other.Y == newY {
			return false // mouvement annulé
		}
	}

	// Appliquer le mouvement
	MoveUnit(unit, moveX, moveY)

	return true
}

func checkEnemiesInRange(unit *types.Unit) []*types.Unit {
	var enemiesInRange []*types.Unit
	for i := range state.Units {
		other := &state.Units[i]
		if IsUnitAlly(unit, other) {
			continue
		}
		dist := manhattanDistance(unit.X, unit.Y, other.X, other.Y)

		if dist <= unit.UnitType.Stats.AttackRange {
			enemiesInRange = append(enemiesInRange, other)
		}
	}
	return enemiesInRange
}

func pickTarget(targets []*types.Unit) *types.Unit {
	unitTypeWeights := map[string]int{
		"cavalry":  2,
		"infantry": 1,
		"archer":   1,
	}

	// Calcul du poids total de toutes les cibles
	totalWeight := 0
	for _, enemy := range targets {
		totalWeight += unitTypeWeights[enemy.UnitType.Name]
	}
	if totalWeight == 0 {
		return targets[0]
	}
	// Tirage aléatoire pondéré
	randomRoll := rand.Intn(totalWeight)
	cumulativeWeight := 0
	for _, enemy := range targets {
		cumulativeWeight += unitTypeWeights[enemy.UnitType.Name]
		if randomRoll < cumulativeWeight {
			return enemy
		}
	}

	// Fallback (ne devrait pas arriver)
	return targets[0]
}
