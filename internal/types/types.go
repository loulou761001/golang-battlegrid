package types

import "battle-sim/assets"

type Tile struct {
	Terrain string // "grass", "dirt", etc.
}

type Stats struct {
	Attack      int
	Health      int
	Defense     int
	AttackRange int
	Speed       int
}

type UnitType struct {
	Name   string
	Symbol string
	Stats  Stats
}

type Unit struct {
	Id          int
	UnitType    UnitType
	Team, X, Y  int
	DamageTaken int
}

type CombatResult struct {
	DamageDealt      int
	Target, Attacker *Unit
}

type CursorPos struct {
	X, Y int
}

var UnitTypes = map[string]UnitType{
	"infantry": {
		Name:   "Infantry",
		Symbol: assets.InfantrySymbol,
		Stats: Stats{
			Health:      6,
			Attack:      4,
			Defense:     5,
			Speed:       1,
			AttackRange: 1,
		},
	},
	"cavalry": {
		Name:   "Cavalry",
		Symbol: assets.CavalrySymbol,
		Stats: Stats{
			Health:      5,
			Attack:      7,
			Defense:     4,
			Speed:       2,
			AttackRange: 1,
		},
	},
	"archer": {
		Name:   "Archer",
		Symbol: assets.ArcherSymbol,
		Stats: Stats{
			Health:      4,
			Attack:      5,
			Defense:     2,
			Speed:       1,
			AttackRange: 3,
		},
	},
}
