package controllers

import (
	"bytes"
	"strings"
	"time"

	"atomicgo.dev/schedule"
	"github.com/pterm/pterm"

	"gplcheck/pkg/common"
)

type StatusViewController struct {
	text      string
	viewWidth int

	buff *bytes.Buffer

	notifier *common.Notifier
	progress *pterm.ProgressbarPrinter
}

func NewStatusViewController(notifier *common.Notifier) *StatusViewController {
	var buf []byte
	c := &StatusViewController{
		notifier: notifier,
		text:     "Ready",
		buff:     bytes.NewBuffer(buf),
	}
	c.run()
	return c
}

func (c *StatusViewController) GetText() string {
	return c.text
}

func (c *StatusViewController) SetWidth(width int) {
	c.viewWidth = width
}

func (c *StatusViewController) Write(p []byte) (n int, err error) {
	if len(p) > 2 {
		c.buff.Truncate(0)
	}
	return c.buff.Write(p)
}

func (c *StatusViewController) run() {
	go func() {
		for {
			select {
			case <-c.notifier.CloseChan():
				return
			case cmd := <-c.notifier.StatusChan():
				switch cmd := cmd.(type) {
				case *common.StatusCommand:
					c.text = cmd.Status
					c.notifier.RefreshChan() <- true
				case *common.ProgressStartCommand:
					pterm.DisableColor()
					c.progress = pterm.DefaultProgressbar.
						WithTotal(cmd.Total).
						WithMaxWidth(c.viewWidth).
						WithTitle(cmd.FileName).
						WithBarFiller(" ").
						WithWriter(c)
					c.progress, _ = c.progress.Start()
					c.text = strings.TrimSpace(c.buff.String())
					c.notifier.RefreshChan() <- true
				case *common.ProgressUpdateCommand:
					c.progress.Add(cmd.Progress)
					c.text = strings.TrimSpace(c.buff.String())
					c.notifier.RefreshChan() <- true
				case *common.ProgressEndCommand:
					_, _ = c.progress.Stop()
					c.text = strings.TrimSpace(c.buff.String())

					c.progress = nil
					c.buff.Reset()

					schedule.After(3*time.Second, func() {
						c.notifier.StatusChan() <- &common.StatusCommand{Status: "Ready"}
					})

					c.notifier.RefreshChan() <- true
				}
			}
		}
	}()
}
