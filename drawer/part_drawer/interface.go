package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
)

type IPartDrawer interface {
	DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int)
}
