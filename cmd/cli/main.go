package main

import (
	"battle-sim/internal/gamelogic"
	"battle-sim/internal/ui"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	gamelogic.GenerateBattlefield()

	if err := ui.PrintGui(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}
