package gingerinterface

import (
	"image"
	"image/color"
)

type IDrawer interface {
	Draw(img *image.RGBA, col color.Color, x1, y1, x2, y2 float64)
}
