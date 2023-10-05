package tui

import (
	"os"
	"path"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type FileView struct {
	tview.TreeView
	root string
}

func NewFileView(dirOrFile string) *FileView {
	if dirOrFile == "" {
		dirOrFile = path.Dir(".")
	}

	if _, err := os.Stat(dirOrFile); err != nil {
		dirOrFile = path.Dir(".")
	}

	root := tview.NewTreeNode(dirOrFile)
	addTreeNode(root, dirOrFile)

	v := &FileView{
		TreeView: *tview.NewTreeView().SetRoot(root).SetCurrentNode(root),
	}
	v.SetBorder(true)
	v.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		node := v.GetCurrentNode()
		dir := len(node.GetChildren()) > 0
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
		}
		return event
	})
	return v
}

func addTreeNode(target *tview.TreeNode, path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		node := tview.NewTreeNode(file.Name()).
			SetReference(filepath.Join(path, file.Name()))
		if file.IsDir() {
			node.SetColor(tcell.ColorBlue).
				Collapse()
			addTreeNode(node, filepath.Join(path, file.Name()))
		}
		target.AddChild(node)
	}
}
