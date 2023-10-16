package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MainFrame struct {
	tview.Flex
	filePanel   *FileView
	resultPanel *ResultView
	status      *StatusView

	app *tview.Application
}

func NewMainFrame(app *tview.Application, file *FileView, result *ResultView, status *StatusView) *MainFrame {
	m := &MainFrame{
		Flex:        *tview.NewFlex(),
		filePanel:   file,
		resultPanel: result,
		status:      status,

		app: app,
	}
	m.DefaultLayout()
	m.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			if m.filePanel.HasFocus() {
				app.SetFocus(m.resultPanel)
			} else {
				app.SetFocus(m.filePanel)
			}
			return nil
		}
		return event
	})
	return m
}

func (m *MainFrame) DefaultLayout() {
	main := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(m.filePanel, 0, 1, true).
		AddItem(m.resultPanel, 0, 3, false)
	m.Flex.Clear().SetDirection(tview.FlexRow).
		AddItem(main, 0, 1, true).
		AddItem(m.status, 1, 1, false)
}
