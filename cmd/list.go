/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tjmblake/worktree-manager/git"
	"github.com/tjmblake/worktree-manager/scanner"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists git worktrees",
	Long: `Lists git worktrees from the current directory.
	Can utilise flags to display additional information.`,
	Run: func(cmd *cobra.Command, args []string) {
		bareDir, _ := os.Getwd()
		gitClient := git.Git{BareDir: bareDir}
		worktrees := gitClient.GetWorktreeList()

		bs := scanner.BatchScanner{BareDir: bareDir}
		bs.Run(&worktrees)

		table, _ := worktrees.Table();
		fmt.Print(table)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
