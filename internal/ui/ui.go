package ui

import (
	"battle-sim/assets"
	"battle-sim/internal/gamelogic"
	"battle-sim/internal/state"
	"battle-sim/internal/types"
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

var battlefield [][]types.Tile
var renderBuffer [][]string
var unitLookup map[[2]int]*types.Unit // O(1) unit lookup

func getUnitName(unit types.Unit) string {
	var nameToPrint string
	if unit.Team == 1 {
		nameToPrint = assets.GreenText + "[" + unit.UnitType.Symbol + "]" + unit.UnitType.Name + assets.ResetText
	} else {
		nameToPrint = assets.RedText + "[" + unit.UnitType.Symbol + "]" + unit.UnitType.Name + assets.ResetText
	}
	return nameToPrint
}

func initBattlefield() {
	battlefield = make([][]types.Tile, state.BattlefieldSize[1])
	renderBuffer = make([][]string, state.BattlefieldSize[1])

	// Fill battlefield with default terrain
	for y := range battlefield {
		battlefield[y] = make([]types.Tile, state.BattlefieldSize[0])
		renderBuffer[y] = make([]string, state.BattlefieldSize[0])
		for x := range battlefield[y] {
			battlefield[y][x].Terrain = "grass" // default terrain
			renderBuffer[y][x] = ""             // empty buffer so everything draws first time
		}
	}
}

func buildUnitLookup() {
	unitLookup = make(map[[2]int]*types.Unit, len(state.Units))
	for i := range state.Units {
		unit := &state.Units[i]
		if gamelogic.IsUnitAlive(unit) {
			unitLookup[[2]int{unit.X, unit.Y}] = unit
		}
	}
}

func getCellRender(x, y int) string {
	// Check if unit is here
	if unit, ok := unitLookup[[2]int{x, y}]; ok {
		if unit.Team == 1 {
			return assets.GreenText + "[" + unit.UnitType.Symbol + "]" + assets.ResetText
		}
		return assets.RedText + "[" + unit.UnitType.Symbol + "]" + assets.ResetText
	}

	// Otherwise, terrain
	switch battlefield[y][x].Terrain {
	case "grass":
		return assets.GrassTexture
	default:
		return assets.GrassTexture
	}
}

func drawBattlefield(v *gocui.View) {
	v.Clear()
	buildUnitLookup() // Update unit positions for this frame

	for y := 0; y < state.BattlefieldSize[1]; y++ {
		var row strings.Builder
		for x := 0; x < state.BattlefieldSize[0]; x++ {
			cell := getCellRender(x, y)
			if renderBuffer[y][x] != cell {
				renderBuffer[y][x] = cell
			}
			row.WriteString(cell)
		}
		fmt.Fprintln(v, row.String())
	}

	// Print status line
	fmt.Fprintf(v, "%d, %d Current player: %d, current turn: %d",
		state.Cursor.X, state.Cursor.Y,
		state.CurrentTurn["activePlayer"],
		state.CurrentTurn["number"],
	)
}

func PrintGui() error {
	initBattlefield()
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

	gamelogic.SetControls(g)
	return g.MainLoop()
}
