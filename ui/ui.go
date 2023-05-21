package ui

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/tjmblake/worktree-manager/scanner"
)

type UI struct {
	tw      *tabwriter.Writer
	bareDir string
}

func CreateUI(bareDir string) UI {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	return UI{tw: writer, bareDir: bareDir}
}

func (ui UI) DisplayScanResults(data []scanner.ScanResponse) {
	fmt.Fprintf(ui.tw, "\n%v\t%v\t%v\t\n", "Index", "Branch", "Modified")

	for i, val := range data {
		since := time.Since(val.Data)
		howLongAgo := fmt.Sprintf("%v Days %v Hours", int(since.Hours())/24, int(since.Hours())%24)
		fmt.Fprintf(ui.tw, "%v\t%v\t%v\t\n", i, val.Worktree.Branch(), howLongAgo)
	}

	ui.tw.Flush()
}

func (ui UI) RequestUserSelection(data []scanner.ScanResponse) []int {
	var selection string

	fmt.Println("\nInput comma-seperated selection to remove: ")
	fmt.Scanln(&selection)

	fmt.Println("\nYou chose to remove: ")
	strSelections := strings.Split(selection, ",")

	var intSelections []int

	for _, str := range strSelections {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("%v cannot be converted to an int, skipping...", str)
			continue
		}
		intSelections = append(intSelections, num)
	}

	for _, index := range intSelections {
		fmt.Fprintf(ui.tw, "%v\t\n", data[index].Worktree.Branch())
	}

	ui.tw.Flush()

	return intSelections
}

func (ui UI) RequestUnsafeWorktreeRemoval(branch string) bool {
	for {
		var selection string
		fmt.Printf("\n\t- %v contains modified / unstaged files.\nAre you sure you would like to remove? (Y|N)\n", branch)
		fmt.Scanln(&selection)

		if selection == "Y" {
			return true
		}

		if selection == "N" {
			return false
		}
	}
}
