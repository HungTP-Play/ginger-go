package ginger

import (
	"log"
	"testing"

	"github.com/HungTP-Play/ginger-go/constant"
)

func BenchmarkDrawIdenticon(t *testing.B) {
	info := "Tom_Ford_Black_Orchid"
	outputDir := "output"
	imageSize := 500
	padding := float64(imageSize) * 0.05
	_, err := DrawIdenticon(info, outputDir, imageSize, int(padding), constant.JPG, false)
	if err != nil {
		log.Printf(err.Error())
	}
}

func BenchmarkDrawGithubIdenticon(t *testing.B) {
	info := "Tom_Ford_Black_Orchid"
	outputDir := "output"
	imageSize := 500
	padding := float64(imageSize) * 0.05
	_, err := DrawGithubIdenticon(info, outputDir, imageSize, int(padding), constant.JPG)
	if err != nil {
		log.Printf(err.Error())
	}
}
