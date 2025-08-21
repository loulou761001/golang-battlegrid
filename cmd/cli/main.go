package main

import (
	"battle-sim/internal/ui"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {

	//gamelogic.GenerateDefaultBattlefield()
	//ui.PrintGui()
	if err := ui.PrintGui(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
