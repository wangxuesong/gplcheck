package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"gplcheck/pkg/common"
)

type ResultViewController struct {
	tview.TableContentReadOnly
	notifier *common.Notifier

	lock   sync.RWMutex
	data   []common.LogEntry
	source []string

	clearHandler func()
}

var (
	headers = []string{"#", "time", "level", "message", "line"}
)

func NewResultViewController(notifier *common.Notifier) *ResultViewController {
	c := &ResultViewController{
		notifier: notifier,

		lock: sync.RWMutex{},
		data: []common.LogEntry{},
	}
	c.run()
	return c
}

func (c *ResultViewController) GetCell(row, column int) *tview.TableCell {
	if row == -1 || column == -1 {
		return nil
	}

	// set table header
	if row == 0 {
		return c.headers(column)
	}
	return c.loadData(row, column)
}

func (c *ResultViewController) GetRowCount() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.data) + 1
}

func (c *ResultViewController) GetColumnCount() int {
	return len(headers)
}

func (c *ResultViewController) headers(col int) *tview.TableCell {
	tc := tview.NewTableCell(" " + headers[col] + " ").
		SetAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorYellow).
		SetBackgroundColor(tcell.ColorBlack).
		SetSelectable(false)
	return tc
}

func (c *ResultViewController) loadData(row int, column int) (tc *tview.TableCell) {
	if column == 0 {
		tc = tview.NewTableCell(strconv.Itoa(row)).
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorYellow).
			SetBackgroundColor(tcell.ColorBlack).
			SetSelectable(true)
	} else {
		tc = c.getData(row, column)
	}
	return tc
}

func (c *ResultViewController) getData(row int, column int) *tview.TableCell {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if row-1 >= len(c.data) {
		return nil
	}
	var value string
	switch column {
	case 1:
		value = c.data[row-1].Time.Format("2006-01-02 15:04:05.000000")
	case 2:
		value = c.data[row-1].Phase
	case 3:
		value = c.data[row-1].Message
	case 4:
		value = fmt.Sprintf("%d", c.data[row-1].Line)
	default:
		value = fmt.Sprintf("$_error on column %d", column)
	}
	tc := tview.NewTableCell(value).
		SetAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorWhite).
		SetSelectable(true)
	return tc
}

func (c *ResultViewController) SetClearHandler(f func()) {
	c.clearHandler = f
}

func (c *ResultViewController) run() {
	go func() {
		for {
			select {
			case <-c.notifier.CloseChan():
				return
			case cmd := <-c.notifier.LogChan():
				switch cmd := cmd.(type) {
				case *common.LogCommand:
					c.lock.Lock()
					c.data = append(c.data, cmd.Entry)
					if cmd.Entry.Line > 0 && cmd.Entry.Line <= len(c.source) {
						line := c.source[cmd.Entry.Line-1]
						if strings.Index(line, "[\"") != 0 {
							c.source[cmd.Entry.Line-1] = fmt.Sprintf(`["%d"]%s[""]`, cmd.Entry.Line, line)
						}
					}
					c.lock.Unlock()
					c.notifier.RefreshChan() <- true
				case *common.ClearCommand:
					c.lock.Lock()
					c.data = []common.LogEntry{}
					c.source = []string{}
					c.lock.Unlock()
					if c.clearHandler != nil {
						c.clearHandler()
					}
					c.notifier.RefreshChan() <- true
				case *common.SourceCommand:
					c.lock.Lock()
					c.source = strings.Split(cmd.Source, "\n")
					c.lock.Unlock()
				}

			}
		}
	}()
}

func (c *ResultViewController) Refresh() {
	go func() { c.notifier.RefreshChan() <- true }()
}

func (c *ResultViewController) GetSource(l int) (string, string) {
	source := c.source
	return strings.Join(source, "\n"), fmt.Sprintf("%d", c.data[l].Line)
}
