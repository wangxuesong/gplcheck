package tui

import (
	"fmt"
	"io"
	"os"

	"github.com/rivo/tview"

	"gplcheck/pkg/common"
	"gplcheck/pkg/worker"
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
					go func() {
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
						// parse file
						w := worker.NewParseWorker(t.notifier)
						script, err := w.Run(string(text))
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
