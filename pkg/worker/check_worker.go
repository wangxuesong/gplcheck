package worker

import (
	"errors"
	"time"

	"github.com/hashicorp/go-multierror"

	"gplcheck/pkg/common"
	"procinspect/pkg/checker"
	"procinspect/pkg/semantic"
)

type CheckWorker struct {
	notifier *common.Notifier
}

func NewCheckWorker(notifier *common.Notifier) *CheckWorker {
	return &CheckWorker{
		notifier: notifier,
	}
}

func (c *CheckWorker) Run(script *semantic.Script) {
	if script == nil {
		return
	}
	v := checker.NewValidVisitor()
	_ = script.Accept(v)

	var errs *multierror.Error
	if errors.As(v.Error(), &errs) && errs != nil {
		for _, e := range errs.Errors {
			err := e.(checker.SqlValidationError)
			logEntry := common.LogEntry{
				Time:    time.Now(),
				Phase:   "check",
				Message: err.Error(),
				Line:    err.Line,
			}
			c.notifier.LogChan() <- &common.LogCommand{Entry: logEntry}
			time.Sleep(100 * time.Millisecond)
		}
	}
}
