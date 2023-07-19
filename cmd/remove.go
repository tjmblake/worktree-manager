/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tjmblake/worktree-manager/git"
	"github.com/tjmblake/worktree-manager/utils"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove worktree/s from your directory.",
	Long: `Remove worktree/s from your directory.

Example: Remove Worktree at Index 1
- wm remove 1 [...] 
- wm rm 1 [...] 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		bareDir, _ := os.Getwd()
		gitClient := git.Git{BareDir: bareDir}
		worktreeList := gitClient.GetWorktreeList()

		selection := utils.AskForWorktreeIndexToDelete()

		for _, worktree := range worktreeList {
			isSelected := false

			for _, index := range selection {
				if worktree.Index == index {
					isSelected = true
				}
			}

			if isSelected {
				gitClient.RemoveWorktree(worktree.Branch, !worktree.SafeToRemove)
				fmt.Println(worktree.Branch, " has been removed!")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
