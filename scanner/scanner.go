package scanner

import (
	"io/fs"
	"log"
	"os"
	"strings"
	"time"
)

var ignoreList []string = []string{"node_modules", ".yarn"}

type Scanner struct {
	Worktree Scannable
	channel  chan ScanResponse
	BareDir  string
}

type ScanResponse struct {
	Worktree Scannable
	Data     time.Time
	BareDir  string
}

func NewScanner(worktree Scannable, bareDir string, channel chan ScanResponse) Scanner {
	return Scanner{Worktree: worktree, channel: channel, BareDir: bareDir}
}

func (s Scanner) ScanLastModified() {
	rootFS := os.DirFS(s.Worktree.Path())
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

	s.channel <- ScanResponse{Worktree: s.Worktree, Data: lastModified, BareDir: s.BareDir}
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
