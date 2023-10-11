package app

import (
	"net/http"
	_ "net/http/pprof"

	"gplcheck/pkg/tui"
)

type App struct {
	Tui *tui.Tui
}

func NewApp(t *tui.Tui) *App {
	a := &App{}
	a.Tui = t
	return a
}

func (a *App) Run() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	a.Tui.Run()
}
