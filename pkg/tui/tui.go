package tui

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rivo/tview"

	"gplcheck/pkg/common"
	"gplcheck/pkg/worker"
	"procinspect/pkg/semantic"
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

	// generate test log data
	go func() {
		// delay 1 second
		time.Sleep(1 * time.Second)
		f, err := os.Open("./test.ast")
		if err != nil {
			return
		}
		defer f.Close()

		gr, err := gzip.NewReader(f)
		if err != nil {
			return
		}
		defer gr.Close()

		// read file to string
		text, err := io.ReadAll(gr)
		if err != nil {
			fmt.Println(err)
			return
		}
		// parse file
		script, err := semantic.NewNodeDecoder[*semantic.Script]().Decode(text)
		w := worker.NewCheckWorker(t.notifier)
		w.Run(script)
	}()

	go func() {
		for {
			select {
			case <-t.notifier.CloseChan():
				return
			case <-t.notifier.RefreshChan():
				t.App.Draw()
			}
		}
	}()

	t.App.Run()
}

func (t *Tui) prepareViews() error {
	t.App.SetRoot(t.Main, true)
	return nil
}
