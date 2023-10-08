package controllers

import (
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type FileViewController struct {
	rootPath string
}

func NewFileViewController() *FileViewController {
	return &FileViewController{rootPath: "."}
}

func (c *FileViewController) GetRootNode() (root *tview.TreeNode) {
	name, _ := filepath.Abs(c.rootPath)
	base := filepath.Base(name)
	root = tview.NewTreeNode(base).SetReference(name)
	addTreeNode(root, name)
	return
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
