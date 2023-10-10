package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"gplcheck/pkg/controllers"
)

type StatusView struct {
	tview.TextView

	controller *controllers.StatusViewController
}

func NewStatusView(c *controllers.StatusViewController) *StatusView {
	return &StatusView{
		TextView:   *tview.NewTextView().SetText("Ready"),
		controller: c,
	}
}

func (v *StatusView) Draw(s tcell.Screen) {
	v.SetText(v.controller.GetText())
	v.TextView.Draw(s)
}
