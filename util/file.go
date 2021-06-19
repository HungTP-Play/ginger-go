package util

import (
	"io/fs"
	"log"
	"os"
)

/// Create directory strucsture for output file if it is not exist
func PrepareOutput(outputPath string) error {
	log.Printf("Checking output dir path: %v", outputPath)
	if _, err := os.Stat(outputPath); os.IsNotExist(err) == true {
		err := os.MkdirAll(outputPath, fs.ModePerm)
		return err
	}
	return nil
}
