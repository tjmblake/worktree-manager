package git

import (
	"log"
	"os/exec"
	"strings"

	"github.com/tjmblake/worktree-manager/models"
)

type Git struct {
	BareDir string
}

func (g Git) GetWorktreeList() models.WorktreeList {
	cmd := exec.Command("git", "worktree", "list", "--porcelain")
	cmd.Dir = g.BareDir
	rawWorktreeList, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	// Remove Header
	postHeader := strings.Index(string(rawWorktreeList), "\nworktree")
	rawWorktreeList = (rawWorktreeList)[postHeader+1:]

	worktrees := parseWorktrees(rawWorktreeList, g.BareDir)

	g.checkSafeToRemove(&worktrees)

	return worktrees
}

func (g Git) CheckWorktreeRemovalSafe(branch string) bool {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = g.BareDir + "/" + branch
	output, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	if len(output) == 0 {
		return true
	}

	return false
}

func (g Git) RemoveWorktree(branch string, force bool) {
	forceArg := ""

	if force {
		forceArg = "--force"
	}

	cmd := exec.Command("git", "worktree", "remove", branch, forceArg)
	cmd.Dir = g.BareDir
	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
}

func parseWorktrees(rawWorktreeList []byte, dir string) models.WorktreeList {
	worktrees := []models.Worktree{}
	rawWorktrees := strings.Split(string(rawWorktreeList), "\n\n")

	for i, v := range rawWorktrees {
		split := strings.Split(v, "\n")
		newTree := models.Worktree{Index: i}

		for _, s := range split {
			val := strings.Split(s, " ")

			if len(val) == 2 {
				if val[0] == "worktree" {
					newTree.Path = val[1]
				}
				if val[0] == "branch" {
					newTree.Branch = strings.TrimPrefix(val[1], "refs/heads/")
				}
			}
		}

		if newTree.Validate() {
			worktrees = append(worktrees, newTree)
		}

	}

	return worktrees
}

func (g Git) checkSafeToRemove(worktrees *models.WorktreeList) {
	for i, worktree := range *worktrees {
		(*worktrees)[i].SafeToRemove = g.CheckWorktreeRemovalSafe(worktree.Branch)
	}
}
