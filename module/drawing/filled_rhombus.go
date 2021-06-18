package drawing

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

type RhombusDrawer struct {
	Name string
}

func (r *RhombusDrawer) Draw(img *image.RGBA, col color.Color, x1, y1, x2, y2 float64) {
	DrawRhombus(img, col, x1, y1, x2, y2)
}

func DrawRhombus(img *image.RGBA, col color.Color, x1, y1, x2, y2 float64) {
	xCenter := x1 + (x2-x1)/2
	yCenter := y1 + (y2-y1)/2

	xTop, yTop := xCenter, y1
	xRight, yRight := x2, yCenter
	xBottom, yBottom := xCenter, y2
	xLeft, yLeft := x1, yCenter

	gc := draw2dimg.NewGraphicContext(img) // Prepare new image context

	gc.SetFillColor(col) // set the color

	gc.MoveTo(xTop, yTop) // top --> right
	gc.LineTo(xRight, yRight)

	gc.MoveTo(xRight, yRight) // right --> bottom
	gc.LineTo(xBottom, yBottom)

	gc.MoveTo(xBottom, yBottom) // bottom --> left
	gc.LineTo(xLeft, yLeft)

	gc.MoveTo(xLeft, yLeft) // left --> top
	gc.LineTo(xTop, yTop)

	// Set the linewidth to zero
	gc.SetLineWidth(0)

	// Fill the stroke so the rectangle will be filled
	gc.FillStroke()
}
