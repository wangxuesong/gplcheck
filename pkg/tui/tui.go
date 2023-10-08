package tui

import (
	"time"

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

	defer close(t.notifier.CloseChan())

	// generate test log data
	go func() {
		// delay 1 second
		time.Sleep(1 * time.Second)
		data := [][]string{
			{"1696160426.41959", "warn", "unsupported: update set multiple columns with select", "6771"},
			{"1696160426.4196968", "warn", "unsupported: update set multiple columns with select", "6771"},
			{"1696160426.41972", "warn", "unsupported: update set multiple columns with select", "6771"},
			{"1696160426.419729", "warn", "unsupported: update set multiple columns with select", "6771"},
			{"1696160426.41959", "warn", "unsupported: update set multiple columns with select", "6771"},
		}
		for _, d := range data {
			d := common.LogEntry{
				Time:    time.Now(),
				Phase:   "check",
				Message: d[2],
				Line:    6771,
			}
			t.notifier.LogChan() <- d
			time.Sleep(1 * time.Second)
		}
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
