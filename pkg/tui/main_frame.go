package tui

import "github.com/rivo/tview"

type MainFrame struct {
	tview.Flex
	filePanel   *tview.TreeView
	resultPanel *tview.Table
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
	m.filePanel = tview.NewTreeView()
	m.filePanel.SetBorder(true)

	m.resultPanel = tview.NewTable()
	m.resultPanel.SetBorder(true)

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
	root := tview.NewTreeNode("Test Data")
	root.AddChild(tview.NewTreeNode("A").AddChild(tview.NewTreeNode("AA")))
	root.AddChild(tview.NewTreeNode("B"))
	m.filePanel.SetRoot(root)

	m.resultPanel.SetTitle("Result")
}
