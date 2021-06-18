package drawing

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

type TriangleDrawer struct {
	Name string
}

func (r *TriangleDrawer) Draw(img *image.RGBA, col color.Color, x1, y1, x2, y2 float64) {
	DrawTriangle(img, col, x1, y1, x2, y2)
}

func DrawTriangle(img *image.RGBA, col color.Color, x1, y1, x2, y2 float64) {
	xTop := x1 + (x2-x1)/2
	gc := draw2dimg.NewGraphicContext(img) // Prepare new image context

	gc.SetFillColor(col) // set the color

	gc.MoveTo(xTop, y1) // top --> bottom right
	gc.LineTo(x2, y2)

	gc.MoveTo(x2, y2) // bottom right --> bottom left
	gc.LineTo(x1, y2)

	gc.MoveTo(x1, y2) // bottom left --> top
	gc.LineTo(xTop, y1)

	// Set the linewidth to zero
	gc.SetLineWidth(0)

	// Fill the stroke so the rectangle will be filled
	gc.FillStroke()
}
