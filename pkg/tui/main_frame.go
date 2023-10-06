package tui

import "github.com/rivo/tview"

type MainFrame struct {
	tview.Flex
	filePanel   *FileView
	resultPanel *ResultView
	status      *tview.TextView
}

func NewMainFrame() *MainFrame {
	m := &MainFrame{
		Flex: *tview.NewFlex(),
	}
	m.makeUIComponents()
	m.DefaultLayout()
	m.loadTestData()
	return m
}

func (m *MainFrame) makeUIComponents() {
	m.filePanel = NewFileView("")

	m.resultPanel = NewResultView()

	m.status = tview.NewTextView().SetText("Ready")
}

func (m *MainFrame) DefaultLayout() {
	main := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(m.filePanel, 0, 1, true).
		AddItem(m.resultPanel, 0, 3, false)
	m.Flex.Clear().SetDirection(tview.FlexRow).
		AddItem(main, 0, 1, true).
		AddItem(m.status, 1, 1, false)
}

func (m *MainFrame) loadTestData() {

	m.resultPanel.SetTitle("Result")
	m.resultPanel.loadTestData()
}
