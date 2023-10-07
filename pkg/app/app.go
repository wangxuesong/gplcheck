package app

import "gplcheck/pkg/tui"

type App struct {
	Tui *tui.Tui
}

func NewApp(t *tui.Tui) *App {
	a := &App{}
	a.Tui = t
	return a
}

func (a *App) Run() {
	a.Tui.Run()
}
