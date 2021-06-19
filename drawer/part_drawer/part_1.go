package partdrawer

import (
	"image"
	"image/color"
)

type Part1Drawer struct {
	Name string
}

func (p *Part1Drawer) DrawPartBitmap(color color.Color, image image.Image) {}
func (p *Part1Drawer) DrawPartSvg(color color.Color, image image.Image)    {}
