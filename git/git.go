package git

import (
	"log"
	"os/exec"
	"strings"
)

type Git struct {
	BareDir string
}

func (g Git) GetWorktreeList() []byte {
	cmd := exec.Command("git", "worktree", "list", "--porcelain")
	cmd.Dir = g.BareDir
	rawWorktreeList, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	trimRawWorktreeList(&rawWorktreeList)
	return rawWorktreeList
}

func trimRawWorktreeList(rawWorktreeList *[]byte) {
	postHeader := strings.Index(string(*rawWorktreeList), "\nworktree")
	*rawWorktreeList = (*rawWorktreeList)[postHeader+1:]
}

func (g Git) CheckWorktreeRemovalSafe(branch string) bool {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = g.BareDir + "/" + branch
	output, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	if len(output) == 0 {
		log.Println("Safe To Remove: ", branch)
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
	output, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(output))
}
