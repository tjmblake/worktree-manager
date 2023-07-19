package models

import (
	"fmt"
	"time"

	"github.com/pterm/pterm"
	"github.com/tjmblake/worktree-manager/utils"
)

type Worktree struct {
	Index        int
	Path         string
	Branch       string
	SafeToRemove bool
	LastModified time.Time
}

func (w Worktree) Validate() bool {
	if w.Path == "" || w.Branch == "" {
		return false
	}
	return true
}

func (w Worktree) Row() []string {
	safeToRemove := utils.RedCross

	if w.SafeToRemove {
		safeToRemove = utils.Check
	}

	return []string{fmt.Sprint(w.Index), w.Branch, w.LastModified.Format(time.DateOnly), safeToRemove}
}

type WorktreeList []Worktree

var header = []string{
	"Index",
	"Branch",
	"Last Modified",
	"Safe to remove",
}

func (list *WorktreeList) Table() (string, error) {
	data := pterm.TableData{header}

	for _, worktree := range *list {
		newRow := []string{worktree.Row()[0], worktree.Row()[1], worktree.Row()[2], worktree.Row()[3]}
		data = append(data, newRow)
	}

	return pterm.DefaultTable.WithHasHeader().WithData(data).Srender()
}
