package tui

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/rivo/tview"

	"gplcheck/pkg/common"
	"gplcheck/pkg/worker"
)

type (
	Tui struct {
		App  *tview.Application
		Main *MainFrame

		ctx context.Context

		notifier *common.Notifier
		pages    *tview.Pages
	}
)

func NewTui(app *tview.Application, main *MainFrame, notifier *common.Notifier) *Tui {
	return &Tui{
		App:   app,
		Main:  main,
		pages: tview.NewPages(),

		notifier: notifier,
	}
}

func InitApp() *tview.Application {
	return tview.NewApplication().EnableMouse(true)
}

func (t *Tui) Run() {
	err := t.prepareViews()
	if err != nil {
		panic(err)
	}

	defer close(t.notifier.CloseChan())

	go func() {
		for {
			select {
			case <-t.notifier.CloseChan():
				return
			case <-t.notifier.RefreshChan():
				t.App.Draw()
			case cmd := <-t.notifier.CommandChan():
				switch c := cmd.(type) {
				case *common.ParseCommand:
					if t.ctx != nil {
						return
					}
					ctx, cancel := context.WithCancel(context.Background())
					t.ctx = ctx
					go func() {
						defer func() {
							cancel()
							t.ctx = nil
						}()
						// open file
						f, err := os.Open(c.FilePath)
						if err != nil {
							return
						}
						defer f.Close()

						// read file to string
						text, err := io.ReadAll(f)
						if err != nil {
							fmt.Println(err)
							return
						}
						t.notifier.LogChan() <- &common.ClearCommand{}
						t.notifier.LogChan() <- &common.SourceCommand{Source: string(text)}
						t.notifier.StatusChan() <- &common.StatusCommand{Status: fmt.Sprintf("Parse %s", c.FilePath)}
						time.Sleep(1 * time.Second)
						// parse file
						t.notifier.StatusChan() <- &common.ProgressStartCommand{
							Total:    100,
							FileName: filepath.Base(c.FilePath),
						}
						w := worker.NewParseWorker(t.notifier)
						script, err := w.Run(string(text))
						t.notifier.StatusChan() <- &common.ProgressEndCommand{}
						if script != nil {
							cw := worker.NewCheckWorker(t.notifier)
							cw.Run(script)
						}
					}()
				}
			}
		}
	}()

	t.App.Run()
}

func (t *Tui) prepareViews() error {
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		flex := tview.NewFlex()
		return flex.
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, 0, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

	t.pages.AddPage("main", t.Main, true, true)

	// show splash screen
	view := NewSplashScreen(t.App)
	t.pages.AddPage("splash", modal(view, view.Width(), view.Height()), true, true)

	// close splash screen
	go func() {
		time.Sleep(5 * time.Second)
		t.pages.RemovePage("splash")
		t.App.Draw()
	}()

	t.App.SetRoot(t.pages, true)
	return nil
}
