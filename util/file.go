package util

import (
	"os"
	"strings"
)

/// Create directory strucsture for output file if it is not exist
func PrepareOutput(outputPath string) error {
	arr := strings.Split(outputPath, "/")
	outputDirPath := strings.Join(arr[:len(arr)-1], "/")
	if _, err := os.Stat(outputDirPath); os.IsNotExist(err) == true {
		err := os.MkdirAll(outputDirPath, os.ModeAppend)
		return err
	}
	return nil
}
