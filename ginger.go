package main

import (
	"fmt"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/HungTP-Play/ginger-go/module/builder"
)

func main() {
	email := "huntp.personal@gmail.com"
	identicon := builder.BuildIdenticon([]byte(email), model.Size4x4, 1024)
	fmt.Printf("IDENTICON %v", identicon.DrawingPoints)
}
