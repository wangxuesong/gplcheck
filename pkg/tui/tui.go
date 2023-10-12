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
						cw := worker.NewCheckWorker(t.notifier)
						cw.Run(script)
					}()
				}
			}
		}
	}()

	t.App.Run()
}

func (t *Tui) prepareViews() error {
	t.App.SetRoot(t.Main, true)
	return nil
}
