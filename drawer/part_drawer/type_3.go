package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Type3Drawer struct {
	Name string
}

func (p *Type3Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int) {
	x2, y2 := sprite.TopLeft.X+sprite.Width, sprite.TopLeft.Y+sprite.Height
	x1, y1 := sprite.TopLeft.X, sprite.TopLeft.Y
	xCenter, yCenter := x1+sprite.Width/2, y1+sprite.Height/2
	gc := draw2dimg.NewGraphicContext(image)
	gc.SetFillColor(color)

	if rotation == 0 {
		// TopCenter -> BottomRight
		gc.MoveTo(xCenter, y1)
		gc.LineTo(x2, y2)

		// BottomRight -> BottomLeft
		gc.LineTo(x1, y2)

		// BottomLeft -> TopCenter
		gc.LineTo(xCenter, y1)
	} else if rotation == 1 {
		// TopLeft -> RightCenter
		gc.MoveTo(x1, y1)
		gc.LineTo(x2, yCenter)

		// RightCenter -> BottomLeft
		gc.LineTo(x1, y2)

		// BottomLeft -> TopLeft
		gc.LineTo(x1, y1)
	} else if rotation == 2 {
		// topLeft -> TopRight
		gc.MoveTo(x1, y1)
		gc.LineTo(x2, y1)

		// TopRight -> BottomCenter
		gc.LineTo(xCenter, y2)

		// BottomCenter -> TopLeft
		gc.LineTo(x1, y1)
	} else {
		// TopRight -> BottomRight
		gc.MoveTo(x2, y1)
		gc.LineTo(x2, y2)

		// BottomRight -> LeftCenter
		gc.LineTo(x1, yCenter)

		// LeftCenter -> TopRight
		gc.LineTo(x2, y1)
	}

	gc.SetLineWidth(0)

	// Fill
	gc.FillStroke()
}
