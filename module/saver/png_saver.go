package saver

import (
	"bufio"
	"image"
	"image/png"
	"os"
)

// SaveToPngFile create and save an image to a file using PNG format
func SaveToPngFile(filePath string, m image.Image) error {
	// Create the file
	f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	// Create Writer from file
	b := bufio.NewWriter(f)
	// Write the image into the buffer
	err = png.Encode(b, m)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}
