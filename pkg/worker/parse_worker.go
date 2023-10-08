package worker

import (
	"gplcheck/pkg/common"
	"procinspect/pkg/checker"
	"procinspect/pkg/semantic"
)

type ParseWorker struct {
	notifier *common.Notifier
}

func NewParseWorker(notifier *common.Notifier) *ParseWorker {
	return &ParseWorker{notifier: notifier}
}

func (c *ParseWorker) Run(text string) (*semantic.Script, error) {
	return checker.LoadScript(string(text))

}
