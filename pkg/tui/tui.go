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

func NewTui() *Tui {
	return &Tui{
		App: tview.NewApplication().EnableMouse(true),
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
	t.Main = NewMainFrame()
	t.App.SetRoot(t.Main, true)
	return nil
}
