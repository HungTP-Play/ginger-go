package partdrawer

import (
	"image"
	"image/color"

	"github.com/HungTP-Play/ginger-go/model"
)

type PartDrawer interface {
	DrawSprite(color color.Color, image image.Image, sprite model.Sprite)
}
