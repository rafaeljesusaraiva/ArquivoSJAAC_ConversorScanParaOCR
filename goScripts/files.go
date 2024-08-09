package goScripts

import (
	"log"
	"os"
)

func CleanupTemporaryFolder(temporaryFolder string) {
	err := os.RemoveAll(temporaryFolder)
	if err != nil {
		log.Fatal(err)
	}
}
