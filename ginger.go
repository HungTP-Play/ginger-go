package main

import (
	"fmt"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/HungTP-Play/ginger-go/module/builder"
	"github.com/HungTP-Play/ginger-go/module/drawing"
)

func main() {
	email := "rhombus"
	identicon := builder.BuildIdenticon([]byte(email), model.Size4x4, 162, true, 128)
	fmt.Printf("IDENTICON %v", identicon.Hash)
	drawing.DrawIdenticon(identicon, "static/", drawing.Rhombus)
}
