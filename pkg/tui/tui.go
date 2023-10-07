package tui

import (
	"github.com/rivo/tview"
)

type (
	Tui struct {
		App  *tview.Application
		Main *MainFrame
	}
)

func NewTui(main *MainFrame) *Tui {
	return &Tui{
		App:  tview.NewApplication().EnableMouse(true),
		Main: main,
	}
}

func (t *Tui) Run() {
	err := t.prepareViews()
	if err != nil {
		panic(err)
	}
	t.App.Run()
}

func (t *Tui) prepareViews() error {
	t.App.SetRoot(t.Main, true)
	return nil
}
