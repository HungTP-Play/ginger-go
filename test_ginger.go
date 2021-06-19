package main

import (
	"log"

	"github.com/HungTP-Play/ginger-go/constant"
	"github.com/HungTP-Play/ginger-go/ginger"
)

func main() {
	info := "Tom_Ford_Black_Orchid"
	outputDir := "output"
	_, err := ginger.DrawIdenticon(info, outputDir, 250, 15, constant.JPG)
	if err != nil {
		log.Printf(err.Error())
	}
}
