package main

import (
	"fmt"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/HungTP-Play/ginger-go/module/builder"
	"github.com/HungTP-Play/ginger-go/module/drawing"
)

func main() {
	email := "hungtp.personal@gmail.com"
	identicon := builder.BuildIdenticon([]byte(email), model.Size7x7, 140)
	fmt.Printf("IDENTICON %v", identicon.Hash)
	drawing.DrawIdenticon(identicon)
}
