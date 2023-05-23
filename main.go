package main

import (
	"os"

	"github.com/tjmblake/worktree-manager/git"
	"github.com/tjmblake/worktree-manager/scanner"
	"github.com/tjmblake/worktree-manager/setup"
	"github.com/tjmblake/worktree-manager/ui"
	"github.com/tjmblake/worktree-manager/worktree"
)

func main() {
	bareDir := setup.Setup()
	gitHandler := git.Git{BareDir: bareDir}
	ui := ui.CreateUI(bareDir)

	unparsedWorktrees := gitHandler.GetWorktreeList()
	worktrees := worktree.ParseWorktrees(unparsedWorktrees, bareDir)

	bs := scanner.BatchScanner{BareDir: bareDir}

	paths := make([]scanner.Scannable, len(worktrees))

	for i, val := range worktrees {
		paths[i] = val
	}

	scanResponses := bs.Run(paths)

	ui.DisplayScanResults(scanResponses)
	selections := ui.RequestUserSelection(scanResponses)

	for i, s := range selections {
		relativeWorktree := worktree.GetRelativeWorktreePath(scanResponses[s].Worktree.Path(), bareDir)
		isSafe := gitHandler.CheckWorktreeRemovalSafe(relativeWorktree)

		if !isSafe {
			approved := ui.RequestUnsafeWorktreeRemoval(scanResponses[s].Worktree.Branch())
			if !approved {
				selections[i] = -1
			}
		}

	}

	for _, s := range selections {
		if s == -1 {
			return
		}

		relativeWorktree := worktree.GetRelativeWorktreePath(scanResponses[s].Worktree.Path(), bareDir)
		gitHandler.RemoveWorktree(relativeWorktree, true)
	}

	os.Exit(0)
}
