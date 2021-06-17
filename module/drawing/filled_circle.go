package drawing

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

func DrawCircle(img *image.RGBA, col color.Color, x1, y1, x2, y2 float64) {
	gc := draw2dimg.NewGraphicContext(img) // Prepare new image context
	xc, yc := x1+float64(x2-x1)/2, y1+float64(y2-y1)/2
	w := (x2 - x1)
	reduceConst := 0.96
	gc.SetFillColor(col) // set the color
	gc.MoveTo(xc, y1)    // move to the topleft in the image
	gc.QuadCurveTo(x1+w*reduceConst, y1+w*(1-reduceConst), x2, yc)
	gc.MoveTo(x2, yc)
	gc.QuadCurveTo(x1+w*reduceConst, y1+w*reduceConst, xc, y2)
	gc.MoveTo(xc, y2)
	gc.QuadCurveTo(x1+w*(1-reduceConst), y1+w*reduceConst, x1, yc)
	gc.MoveTo(x1, yc)
	gc.QuadCurveTo(x1+w*(1-reduceConst), y1+w*(1-reduceConst), xc, y1)
	gc.MoveTo(xc, y1)
	gc.SetLineWidth(0)
	// Fill the stroke so the rectangle will be filled
	gc.FillStroke()
}
