package drawing

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

func DrawRect(img *image.RGBA, col color.Color, x1, y1, x2, y2 float64) {
	gc := draw2dimg.NewGraphicContext(img) // Prepare new image context
	gc.SetFillColor(col)                   // set the color
	gc.MoveTo(x1, y1)                      // move to the topleft in the image
	// Draw the lines for the dimensions
	gc.LineTo(x1, y1)
	gc.LineTo(x1, y2)
	gc.MoveTo(x2, y1) // move to the right in the image
	// Draw the lines for the dimensions
	gc.LineTo(x2, y1)
	gc.LineTo(x2, y2)
	// Set the linewidth to zero
	gc.SetLineWidth(0)
	// Fill the stroke so the rectangle will be filled
	gc.FillStroke()
}
