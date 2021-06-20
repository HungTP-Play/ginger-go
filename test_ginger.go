package main

import (
	"log"

	"github.com/HungTP-Play/ginger-go/constant"
	"github.com/HungTP-Play/ginger-go/ginger"
)

func main() {
	info := "🦄🦄🦄🦄🦄"
	outputDir := "output"
	imageSize := 500
	padding := float64(imageSize) * 0.05
	_, err := ginger.DrawGithubIdenticon(info, outputDir, imageSize, int(padding), constant.PNG)
	if err != nil {
		log.Printf(err.Error())
	}
}
