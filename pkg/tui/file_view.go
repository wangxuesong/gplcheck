package tui

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"gplcheck/pkg/common"
	"gplcheck/pkg/controllers"
)

type FileView struct {
	tview.TreeView
	controller *controllers.FileViewController
	notifier   *common.Notifier
}

func NewFileView(notifier *common.Notifier, c *controllers.FileViewController) *FileView {
	root := c.GetRootNode()

	v := &FileView{
		TreeView:   *tview.NewTreeView().SetRoot(root).SetCurrentNode(root),
		controller: c,
		notifier:   notifier,
	}
	v.SetBorder(true)
	v.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		node := v.GetCurrentNode()
		name := node.GetReference().(string)
		info, _ := os.Stat(name)
		dir := info.IsDir()
		switch event.Key() {
		case tcell.KeyRight:
			if dir && !node.IsExpanded() {
				node.Expand()
			}
			return nil
		case tcell.KeyLeft:
			if dir && node.IsExpanded() {
				node.Collapse()
			}
			return nil
		case tcell.KeyEnter:
			if !dir {
				cmd := &common.ParseCommand{FilePath: name}
				v.notifier.CommandChan() <- cmd
				return nil
			}
		}
		return event
	})
	return v
}
