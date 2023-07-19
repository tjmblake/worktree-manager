package scanner

import (
	"io/fs"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tjmblake/worktree-manager/models"
)

var ignoreList []string = []string{"node_modules", ".yarn"}

type WorktreeScanner struct {
	Worktree models.Worktree
	channel  chan models.Worktree
	BareDir  string
}

func NewScanner(worktree models.Worktree, bareDir string, channel chan models.Worktree) WorktreeScanner {
	return WorktreeScanner{Worktree: worktree, channel: channel, BareDir: bareDir}
}

func (s WorktreeScanner) ScanLastModified() {
	rootFS := os.DirFS(s.Worktree.Path)
	var lastModified time.Time

	fs.WalkDir(rootFS, ".", func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if !checkIfIgnored(filePath) {
			return nil
		}

		info, err := d.Info()

		if err != nil {
			log.Fatal(err)
		}

		modTime := info.ModTime()

		if lastModified.Unix() < modTime.Unix() {
			lastModified = modTime
		}

		return nil
	})

	s.Worktree.LastModified = lastModified
	s.channel <- s.Worktree
}

func checkIfIgnored(path string) bool {
	if path == "." {
		return false
	}
	for _, ignore := range ignoreList {
		if strings.HasPrefix(path, ignore) {
			return false
		}
	}
	return true
}
