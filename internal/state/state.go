package state

import (
	"battle-sim/internal/types"
)

var Cursor = types.CursorPos{}

var BattlefieldSize = [2]int{32, 24}

var CurrentTurn = map[string]int{
	"activePlayer": 1,
	"number":       1,
}
var Units = []types.Unit{
	{
		0,
		types.UnitTypes["infantry"],
		1,
		12,
		3,
		0,
	},
	{
		1,
		types.UnitTypes["infantry"],
		1,
		15,
		3,
		0,
	},
	{
		2,
		types.UnitTypes["infantry"],
		1,
		17,
		3,
		0,
	},
	{
		7,
		types.UnitTypes["cavalry"],
		1,
		6,
		2,
		0,
	},

	{
		3,
		types.UnitTypes["infantry"],
		2,
		13,
		20,
		0,
	},
	{
		4,
		types.UnitTypes["infantry"],
		2,
		20,
		20,
		0,
	},
	{
		5,
		types.UnitTypes["infantry"],
		2,
		16,
		20,
		0,
	},
	{
		6,
		types.UnitTypes["archer"],
		2,
		17,
		20,
		0,
	},
}

var LastCombatResult []*types.CombatResult

var WinningTeam = 0
