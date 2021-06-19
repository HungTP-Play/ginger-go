package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
)

type Type16Drawer struct {
	Name string
}

func (p *Type16Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int) {
	// !NOTE: Do nothing here
}
