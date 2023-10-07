package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"gplcheck/pkg/controllers"
)

type ResultView struct {
	tview.Table
	controller *controllers.ResultViewController
}

func NewResultView(c *controllers.ResultViewController) *ResultView {
	v := &ResultView{
		Table: *tview.NewTable().
			SetSelectable(true, false).
			SetFixed(1, 1).
			SetSeparator(tview.Borders.Vertical),
		controller: c,
	}
	v.SetBorder(true)
	v.SetBackgroundColor(tcell.ColorBlack)
	v.SetContent(v.controller)
	return v
}
