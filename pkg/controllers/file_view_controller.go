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
		fullname := filepath.Join(path, file.Name())
		node := tview.NewTreeNode(file.Name()).
			SetReference(fullname)
		info, err := os.Lstat(fullname)

		if err == nil {
			if info.Mode()&os.ModeSymlink == os.ModeSymlink {
				// Handling symlinks. Note, may need to handle recursive symlinks with care to prevent infinite recursion
				target, err := os.Readlink(fullname)
				if err == nil {
					info, err = os.Stat(target)
				}
				if err == nil && info.IsDir() {
					node.SetColor(tcell.ColorBlue).
						Collapse()
					addTreeNode(node, fullname)
				}
			}

		}
		if file.IsDir() {
			node.SetColor(tcell.ColorBlue).
				Collapse()
			addTreeNode(node, fullname)
		}
		target.AddChild(node)
	}
}
