package tui

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"gplcheck/pkg/controllers"
)

type ResultView struct {
	tview.Flex
	table  tview.Table
	source tview.TextView

	controller *controllers.ResultViewController
}

func NewResultView(c *controllers.ResultViewController) *ResultView {
	v := &ResultView{
		Flex: *tview.NewFlex(),
		table: *tview.NewTable().
			SetSelectable(true, false).
			SetFixed(1, 1).
			SetSeparator(tview.Borders.Vertical),
		source: *tview.NewTextView().
			SetDynamicColors(true).
			SetWrap(false).
			SetScrollable(true).
			SetRegions(true),
		controller: c,
	}
	v.SetBorder(true)
	v.table.SetBackgroundColor(tcell.ColorBlack)
	v.table.SetContent(v.controller)
	v.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			l, _ := v.table.GetSelection()
			text, tag := v.controller.GetSource(l - 1)
			v.showSource(true)
			v.controller.Refresh()
			v.source.SetText(text)
			go func() { v.scrollToError(tag) }()
			return nil
		}
		return event
	})

	v.source.SetText("source").SetDisabled(true)

	v.SetDirection(tview.FlexRow).
		AddItem(&v.table, 0, 1, true)

	v.controller.SetClearHandler(func() {
		v.showSource(false)
	})

	return v
}

func (v *ResultView) showSource(show bool) {
	if show {
		if v.GetItemCount() == 1 {
			v.AddItem(&v.source, 0, 1, false)
		}
	} else {
		if v.GetItemCount() == 2 {
			v.RemoveItem(&v.source)
		}
	}
}

func (v *ResultView) scrollToError(tag string) {
	time.Sleep(10 * time.Millisecond)
	v.source.Highlight(tag).ScrollToHighlight()
	v.controller.Refresh()
}
