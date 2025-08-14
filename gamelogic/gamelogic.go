package gamelogic

import (
	"battle-sim/state"
	"battle-sim/types"
	"math/rand/v2"
)

func IsUnitAlive(unit *types.Unit) bool {
	return unit.DamageTaken < unit.UnitType.Stats.Health
}

func IsUnitAlly(unit, other *types.Unit) bool {
	return other.Id == unit.Id || other.Team == unit.Team || !IsUnitAlive(other)
}

// GAME GENERATION ---

func findUnitsOnLine(y int) []types.Unit {
	var result []types.Unit
	var livingUnits []types.Unit
	for _, unit := range state.Units {
		if IsUnitAlive(&unit) {
			livingUnits = append(livingUnits, unit)
		}
	}
	for _, u := range livingUnits {
		if u.Y == y {
			result = append(result, u)
		}
	}
	return result
}

func GenerateBattlefield() [][]string {
	lines := make([][]string, 0, state.BattlefieldSize[1])
	for currentLine := 0; currentLine < state.BattlefieldSize[1]; currentLine++ {
		line := make([]string, state.BattlefieldSize[0])
		var unitsOnLine = findUnitsOnLine(currentLine)
		for _, unit := range unitsOnLine {
			var symbolToDisplay string
			if unit.Team == 1 {
				symbolToDisplay = "\033[32m" + unit.UnitType.Symbol + "\033[0m"
			} else {
				symbolToDisplay = "\033[31m" + unit.UnitType.Symbol + "\033[0m"

			}
			line[unit.X] = symbolToDisplay
		}
		lines = append(lines, line)
	}
	state.Battlefield = lines
	return lines
}

// ---

func MoveUnit(unit *types.Unit, x int, y int) {
	if unit.X+x >= 0 {
		unit.X += x
	}
	if unit.Y+y >= 0 {
		unit.Y += y
	}
}

func resolveCombat(unit *types.Unit, target *types.Unit) types.CombatResult {
	attackerDice := throwDice(6)
	defenderDice := throwDice(6)
	totalAttack := unit.UnitType.Stats.Attack + attackerDice
	totalDefense := target.UnitType.Stats.Defense + defenderDice
	damage := 1

	if defenderDice == 1 || attackerDice == 6 {
		damage = 2
	}

	if totalAttack > totalDefense {
		target.DamageTaken += damage
	}
	results := types.CombatResult{DamageDealt: damage, Target: target, Attacker: unit}
	state.LastCombatResult = append(state.LastCombatResult, &results)
	return results
}

//----

func throwDice(faces int) int {
	// -1 and +1 to avoid 0
	return rand.IntN(faces-1) + 1
}

func changeTurn() {
	if state.CurrentTurn["activePlayer"] == 1 {
		state.CurrentTurn["activePlayer"] = 2
	} else {
		state.CurrentTurn["number"] += 1
		state.CurrentTurn["activePlayer"] = 1
	}
}

func checkWinner() {
	remainingTeam1Units := 0
	remainingTeam2Units := 0
	for i := range state.Units {
		currentUnit := &state.Units[i]
		if IsUnitAlive(currentUnit) {
			if currentUnit.Team == 1 {
				remainingTeam1Units += 1
			} else {
				remainingTeam2Units += 1
			}
		}
	}
	if remainingTeam1Units == 0 {
		state.WinningTeam = 1
	} else if remainingTeam2Units == 0 {
		state.WinningTeam = 2
	}
}

func EndTurn() {
	checkWinner()
	if state.WinningTeam != 0 {
		return
	}
	state.LastCombatResult = []*types.CombatResult{}
	for i := range state.Units {
		currentUnit := &state.Units[i]
		if currentUnit.Team == state.CurrentTurn["activePlayer"] && IsUnitAlive(currentUnit) {
			inRange := checkEnemiesInRange(currentUnit)
			if len(inRange) == 0 {
				moveUnitTowardsClosest(currentUnit)
			} else {
				target := pickTarget(inRange)
				resolveCombat(currentUnit, target)
			}
		}
	}
	changeTurn()
	GenerateBattlefield()
}

// ---

func MoveCursor(moveX, moveY int) {
	maxY := len(state.Battlefield)
	maxX := len(state.Battlefield[0])
	if state.Cursor.X+moveX >= 0 && state.Cursor.X+moveX < maxX {
		state.Cursor.X += moveX
	}
	if state.Cursor.Y+moveY >= 0 && state.Cursor.Y+moveY < maxY {
		state.Cursor.Y += moveY
	}
}
