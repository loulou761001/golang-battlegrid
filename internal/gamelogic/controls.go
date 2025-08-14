package gamelogic

import (
	"log"

	"github.com/jroimartin/gocui"
)

func SetControls(gui *gocui.Gui) {

	var keybindings = []struct {
		key     gocui.Key
		handler func(*gocui.Gui, *gocui.View) error
	}{
		{gocui.KeyCtrlC, func(g *gocui.Gui, v *gocui.View) error { return gocui.ErrQuit }},
		{gocui.KeyArrowUp, func(g *gocui.Gui, v *gocui.View) error {
			MoveCursor(0, -1)
			return nil
		}},
		{gocui.KeyArrowDown, func(g *gocui.Gui, v *gocui.View) error {
			MoveCursor(0, 1)
			return nil
		}},
		{gocui.KeyArrowLeft, func(g *gocui.Gui, v *gocui.View) error {
			MoveCursor(-1, 0)
			return nil
		}},
		{gocui.KeyArrowRight, func(g *gocui.Gui, v *gocui.View) error {
			MoveCursor(1, 0)
			return nil
		}},

		{gocui.KeySpace, func(g *gocui.Gui, v *gocui.View) error {
			EndTurn()
			return nil
		}},
	}

	for _, kb := range keybindings {
		if err := gui.SetKeybinding("", kb.key, gocui.ModNone, kb.handler); err != nil {
			log.Panicln(err)
		}
	}
}
