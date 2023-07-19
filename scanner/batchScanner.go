package scanner

import (
	"github.com/tjmblake/worktree-manager/models"
)


type BatchScanner struct {
	BareDir string
}

func (b BatchScanner) Run(worktreeList *models.WorktreeList) {
	scannerChannel := make(chan models.Worktree)

	for _, lo := range *worktreeList {
		scanner := NewScanner(lo, b.BareDir, scannerChannel)
		go scanner.ScanLastModified()
	}

	for i := 0; i < len(*worktreeList); i++ {
		updatedWorktree := <-scannerChannel
		(*worktreeList)[updatedWorktree.Index] = updatedWorktree
	}

}
