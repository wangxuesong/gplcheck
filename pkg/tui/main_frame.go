package tui

import "github.com/rivo/tview"

type MainFrame struct {
	tview.Flex
	filePanel   *FileView
	resultPanel *ResultView
	status      *StatusView
}

func NewMainFrame(file *FileView, result *ResultView, status *StatusView) *MainFrame {
	m := &MainFrame{
		Flex:        *tview.NewFlex(),
		filePanel:   file,
		resultPanel: result,
		status:      status,
	}
	m.DefaultLayout()
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
