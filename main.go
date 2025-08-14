package main

import (
	"battle-sim/gamelogic"
	"battle-sim/ui"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	gamelogic.GenerateBattlefield()

	if err := ui.PrintGui(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}
