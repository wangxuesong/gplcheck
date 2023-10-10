package controllers

import "gplcheck/pkg/common"

type StatusViewController struct {
	text string

	notifier *common.Notifier
}

func NewStatusViewController(notifier *common.Notifier) *StatusViewController {
	c := &StatusViewController{notifier: notifier, text: "Ready"}
	c.run()
	return c
}

func (c *StatusViewController) GetText() string {
	return c.text
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
				}
			}
		}
	}()
}
