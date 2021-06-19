package partdrawer

import (
	"image"
	"image/color"
)

type PartDrawer interface {
	DrawPartBitmap(color color.Color, image image.Image, )
	DrawPartSvg(color color.Color, image image.Image)
}
