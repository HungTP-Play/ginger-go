package main

import (
	"log"

	"github.com/HungTP-Play/ginger-go/constant"
	"github.com/HungTP-Play/ginger-go/ginger"
)

func main() {
	info := "Identicon_Vesace_Eros"
	outputDir := "output"
	imageSize := 500
	padding := float64(imageSize) * 0.05
	_, err := ginger.DrawIdenticon(info, outputDir, imageSize, int(padding), constant.PNG, true)
	if err != nil {
		log.Printf(err.Error())
	}
}
