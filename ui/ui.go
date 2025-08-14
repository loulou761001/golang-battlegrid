package ui

import (
	"battle-sim/assets"
	"battle-sim/gamelogic"
	"battle-sim/state"
	"battle-sim/types"
	"fmt"
	"strconv"

	"github.com/jroimartin/gocui"
)

func getUnitName(unit types.Unit) string {
	var nameToPrint string
	if unit.Team == 1 {
		nameToPrint = "\033[32m[" + unit.UnitType.Symbol + "] " + unit.UnitType.Name + "\033[0m"
	} else {
		nameToPrint = "\033[31m[" + unit.UnitType.Symbol + "] " + unit.UnitType.Name + "\033[0m"
	}
	return nameToPrint
}

func printBattlefield(v *gocui.View) {
	for y, line := range state.Battlefield {
		for x, cell := range line {
			var cellToPrint string
			if len(cell) == 0 {
				cellToPrint = " "
			} else {
				cellToPrint = cell
			}
			if x == state.Cursor.X && y == state.Cursor.Y {
				// Curseur : fond blanc et texte noir
				fmt.Fprint(v, "\033[47;30m["+cellToPrint+"]\033[0m")
			} else {
				fmt.Fprint(v, "["+cellToPrint+"]")
			}
		}
		fmt.Fprintln(v)
	}
	fmt.Fprint(v, strconv.Itoa(state.Cursor.X)+", "+strconv.Itoa(state.Cursor.Y))
	fmt.Fprint(v, " Current player: "+strconv.Itoa(state.CurrentTurn["activePlayer"])+", current turn: "+strconv.Itoa(state.CurrentTurn["number"]))
}

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
		fmt.Fprint(v, "Aucune unité à cette position")
	}

}

func printCombatInfo(v *gocui.View) {
	if state.WinningTeam != 0 {
		str := fmt.Sprintf("Victory for team %d!", state.WinningTeam)
		fmt.Fprintln(v, str)
		return
	}

	if state.LastCombatResult == nil || len(state.LastCombatResult) == 0 {
		fmt.Fprintln(v, assets.GameLogo)
		return
	}
	for _, result := range state.LastCombatResult {
		attackerName := getUnitName(*result.Attacker)
		targetName := getUnitName(*result.Target)
		str := fmt.Sprintf("%s attacked %s, %d damage dealt.", attackerName, targetName, result.DamageDealt)
		fmt.Fprintln(v, str)
	}
}

func drawBattlefield(v *gocui.View) {
	v.Clear()
	printBattlefield(v)
}

func drawUnitInfo(v *gocui.View) {
	v.Clear()
	printUnitInfo(v)
}

func drawCombatInfo(v *gocui.View) {
	v.Clear()
	printCombatInfo(v)
}

func PrintGui() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	defer g.Close()

	g.SetManagerFunc(func(g *gocui.Gui) error {
		maxX, maxY := g.Size()
		combatHeight := 8
		battleView, _ := g.SetView("battlefield", 0, 0, maxX-1, maxY-combatHeight-1)
		combatView, _ := g.SetView("combatView", 0, maxY-combatHeight, maxX-1, maxY-1)
		infoView, _ := g.SetView("unitInfo", maxX-36, 0, maxX-1, 8)

		drawBattlefield(battleView)
		drawUnitInfo(infoView)
		drawCombatInfo(combatView)
		return nil
	})

	// Définir les keybindings ici (Ctrl+C + flèches)
	gamelogic.SetControls(g)

	return g.MainLoop()
}
