package partdrawer

import (
	"image"
	"image/color"

	"github.com/HungTP-Play/ginger-go/model"
)

type Part1Drawer struct {
	Name string
}

func (p *Part1Drawer) DrawSprite(color color.Color, image image.Image, sprite model.Sprite) {}
