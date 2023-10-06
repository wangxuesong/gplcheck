package controllers

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ResultViewController struct {
	tview.TableContentReadOnly
}

var (
	headers = []string{"#", "time", "level", "message", "line"}
	data    = [][]string{
		{"1696160426.41959", "warn", "unsupported: update set multiple columns with select", "6771"},
		{"1696160426.41959", "warn", "unsupported: update set multiple columns with select", "6771"},
		{"1696160426.41959", "warn", "unsupported: update set multiple columns with select", "6771"},
		{"1696160426.41959", "warn", "unsupported: update set multiple columns with select", "6771"},
		{"1696160426.41959", "warn", "unsupported: update set multiple columns with select", "6771"},
	}
)

func NewResultViewController() *ResultViewController {
	return &ResultViewController{}
}

func (c *ResultViewController) GetCell(row, column int) *tview.TableCell {
	if row == -1 || column == -1 {
		return nil
	}

	// set table header
	if row == 0 {
		return c.loadTestDataHeaders(column)
	}
	return c.loadTestData(row, column)
}

func (c *ResultViewController) GetRowCount() int {
	return len(data) + 1
}

func (c *ResultViewController) GetColumnCount() int {
	return len(headers)
}

func (c *ResultViewController) loadTestDataHeaders(col int) *tview.TableCell {
	tc := tview.NewTableCell(" " + headers[col] + " ").
		SetAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorYellow).
		SetBackgroundColor(tcell.ColorBlack).
		SetSelectable(false)
	return tc
}

func (c *ResultViewController) loadTestData(row int, column int) (tc *tview.TableCell) {
	if column == 0 {
		tc = tview.NewTableCell(strconv.Itoa(row)).
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorYellow).
			SetBackgroundColor(tcell.ColorBlack).
			SetSelectable(true)
	} else {
		tc = tview.NewTableCell(data[row-1][column-1]).
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorWhite).
			SetSelectable(true)
	}
	return tc
}
