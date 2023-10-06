package tui

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ResultView struct {
	tview.Table
}

func NewResultView() *ResultView {
	v := &ResultView{
		Table: *tview.NewTable().
			SetSelectable(true, false).
			SetFixed(1, 1).
			SetSeparator(tview.Borders.Vertical),
	}
	v.SetBorder(true)
	return v
}

func (v *ResultView) loadTestData() {
	// set table header
	headers := []string{"#", "time", "level", "message", "line"}
	for i, header := range headers {
		v.SetCell(0, i, tview.NewTableCell(" "+header+" ").
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorYellow).
			SetBackgroundColor(tcell.ColorBlack).
			SetSelectable(false))
	}

	// set table data
	data := [][]string{
		{"1696160426.41959", "warn", "unsupported: update set multiple columns with select", "6771"},
		{"1696160426.41959", "warn", "unsupported: update set multiple columns with select", "6771"},
		{"1696160426.41959", "warn", "unsupported: update set multiple columns with select", "6771"},
		{"1696160426.41959", "warn", "unsupported: update set multiple columns with select", "6771"},
		{"1696160426.41959", "warn", "unsupported: update set multiple columns with select", "6771"},
	}
	for i, d := range data {
		v.SetCell(i+1, 0, tview.NewTableCell(strconv.Itoa(i+1)).
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorYellow).
			SetBackgroundColor(tcell.ColorBlack).
			SetSelectable(true))
		for j, value := range d {
			v.SetCell(i+1, j+1, tview.NewTableCell(" "+value+" ").
				SetAlign(tview.AlignCenter).
				SetTextColor(tcell.ColorWhite).
				SetSelectable(true))
		}
	}
}
