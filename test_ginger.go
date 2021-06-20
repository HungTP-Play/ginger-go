package main

import (
	"log"

	"github.com/HungTP-Play/ginger-go/constant"
	"github.com/HungTP-Play/ginger-go/ginger"
)

func main() {
	info := "Identicon_Clue_De_Nuit_Pour_Homme"
	outputDir := "output"
	_, err := ginger.DrawIdenticon(info, outputDir, 250, 15, constant.PNG)
	if err != nil {
		log.Printf(err.Error())
	}
}
