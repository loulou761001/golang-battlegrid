package ui

import (
	"battle-sim/assets"
	"battle-sim/internal/state"
	"fmt"

	"github.com/jroimartin/gocui"
)

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

func drawCombatInfo(v *gocui.View) {
	v.Clear()
	printCombatInfo(v)
}
