package worktree

import (
	"strings"
)

type Worktree struct {
	path   string
	branch string
}

func (w Worktree) Path() string {
	return w.path
}

func (w Worktree) Branch() string {
	return w.branch
}

func (w Worktree) validate() bool {
	if w.path == "" || w.branch == "" {
		return false
	}
	return true
}

func parseRawWorktreeList(rawWorktreeList []byte, dir string) []Worktree {
	worktrees := []Worktree{}
	rawWorktrees := strings.Split(string(rawWorktreeList), "\n\n")

	for _, v := range rawWorktrees {
		split := strings.Split(v, "\n")
		newTree := Worktree{}

		for _, s := range split {
			val := strings.Split(s, " ")

			if len(val) == 2 {
				if val[0] == "worktree" {
					newTree.path = val[1]
				}
				if val[0] == "branch" {
					newTree.branch = strings.TrimPrefix(val[1], "refs/heads/")
				}
			}
		}

		if newTree.validate() {
			worktrees = append(worktrees, newTree)
		}
	}

	return worktrees
}

func ParseWorktrees(rawWorktreeData []byte, dir string) []Worktree {
	parsedWorktrees := parseRawWorktreeList(rawWorktreeData, dir)
	return parsedWorktrees
}

func GetRelativeWorktreePath(path string, bareDir string) string {
	return strings.TrimPrefix(path, bareDir+"/")
}
