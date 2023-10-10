package worker

import (
	"errors"
	"time"

	"gplcheck/pkg/common"
	"procinspect/pkg/checker"
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
	script, err := checker.LoadScript(string(text))
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
