package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func AskForWorktreeIndexToDelete() []int {
	var strSelections []string
	var intSelections []int

	for len(intSelections) == 0 {
		selection := ""
		fmt.Printf("\n%v Input comma-seperated selection to remove: ", Question)
		fmt.Scanln(&selection)

		fmt.Println("\nYou chose to remove: ", selection)
		strSelections = strings.Split(selection, ",")

		for _, str := range strSelections {
			if str == "" {
				continue
			}
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("\n%v cannot be converted to an int, skipping...", str)
				continue
			}
			intSelections = append(intSelections, num)
		}
	}

	return intSelections
}
