package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"gplcheck/pkg/tui"
)

func main() {
	// splash screen
	app := tview.NewApplication().EnableMouse(true)
	view := tui.NewSplashScreen(app)

	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, 0, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

	background := tview.NewTextView().
		SetTextColor(tcell.ColorBlue).
		SetText(strings.Repeat("background ", 1000))

	pages := tview.NewPages().
		AddPage("background", background, true, true).
		AddPage("modal", modal(view, view.Width(), view.Height()), true, true)

	app.SetRoot(pages, true)
	app.Run()
}
