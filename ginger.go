package main

import (
	"fmt"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/HungTP-Play/ginger-go/module/builder"
	"github.com/HungTP-Play/ginger-go/module/drawing"
)

func main() {
	email := "stoicmeme@gmail.com"
	identicon := builder.BuildIdenticon([]byte(email), model.Size5x5, 250, false, 128)
	fmt.Printf("IDENTICON %v", identicon.Hash)
	drawing.DrawIdenticon(identicon, "output", drawing.Circle)
}
