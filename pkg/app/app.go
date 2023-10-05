package app

import "gplcheck/pkg/tui"

type App struct {
	Tui *tui.Tui
}

func NewApp() *App {
	a := &App{}
	t := tui.NewTui()
	a.Tui = t
	return a
}

func (a *App) Run() {
	a.Tui.Run()
}
