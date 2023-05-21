package setup

import (
	"log"
	"os"
)

func Setup() string {
	bareDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return bareDir
}
