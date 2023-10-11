package worker

import (
	"errors"
	"strings"
	"time"

	"gplcheck/pkg/common"
	"gplcheck/pkg/utils"
	"procinspect/pkg/parser"
	"procinspect/pkg/semantic"
)

type ParseWorker struct {
	notifier *common.Notifier
}

func NewParseWorker(notifier *common.Notifier) *ParseWorker {
	return &ParseWorker{notifier: notifier}
}

func (c *ParseWorker) Run(text string) (*semantic.Script, error) {
	script, err := c.parse(text)
	if err != nil {
		joinErr, ok := err.(interface{ Unwrap() []error })
		if ok {
			errs := joinErr.Unwrap()
			for _, e := range errs {
				var pe parser.ParseError
				if errors.As(e, &pe) {
					logEntry := common.LogEntry{
						Time:    time.Now(),
						Phase:   "semantic",
						Message: pe.Msg,
						Line:    pe.Line,
					}
					c.notifier.LogChan() <- &common.LogCommand{Entry: logEntry}
				}
			}
		}
	}
	return script, err
}

func (c *ParseWorker) parse(text string) (*semantic.Script, error) {
	t := len(strings.Split(text, "\n"))
	c.notifier.StatusChan() <- &common.ProgressUpdateCommand{Progress: 0, Total: t}
	parser := utils.NewParallelParser(text).
		WithUpdateHandler(func(delta, total int) {
			c.notifier.StatusChan() <- &common.ProgressUpdateCommand{Progress: delta, Total: total}
		})

	return parser.Parse()
}
