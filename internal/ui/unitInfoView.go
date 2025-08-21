package ui

import (
	"battle-sim/internal/gamelogic"
	"battle-sim/internal/state"
	"battle-sim/internal/types"
	"fmt"
	"strconv"

	"github.com/jroimartin/gocui"
)

func printUnitInfo(v *gocui.View) {
	var hoveredUnit types.Unit

	found := false

	for _, unit := range state.Units {
		if unit.X == state.Cursor.X && unit.Y == state.Cursor.Y && gamelogic.IsUnitAlive(&unit) {
			hoveredUnit = unit
			found = true
			break
		}
	}
	if found {
		var nameToPrint = getUnitName(hoveredUnit)
		linesToPrint := []string{
			nameToPrint,
			"ATK: " + strconv.Itoa(hoveredUnit.UnitType.Stats.Attack),
			"DEF: " + strconv.Itoa(hoveredUnit.UnitType.Stats.Defense),
			"HP: " + strconv.Itoa(hoveredUnit.UnitType.Stats.Health-hoveredUnit.DamageTaken) + "/" + strconv.Itoa(hoveredUnit.UnitType.Stats.Health),
		}
		for _, line := range linesToPrint {
			fmt.Fprintln(v, line)
		}
	} else {
		fmt.Fprint(v, "No unit at this position")
	}
}

func drawUnitInfo(v *gocui.View) {
	v.Clear()
	printUnitInfo(v)
}