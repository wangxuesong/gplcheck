package tui

import (
	"github.com/rivo/tview"

	"gplcheck/pkg/common"
)

type (
	Tui struct {
		App  *tview.Application
		Main *MainFrame

		notifier *common.Notifier
	}
)

func NewTui(main *MainFrame, notifier *common.Notifier) *Tui {
	return &Tui{
		App:  tview.NewApplication().EnableMouse(true),
		Main: main,

		notifier: notifier,
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
