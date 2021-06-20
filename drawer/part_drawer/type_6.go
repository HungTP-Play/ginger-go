package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Type6Drawer struct {
	Name string
}

func (p *Type6Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int) {
	x2, y2 := sprite.TopLeft.X+sprite.Width, sprite.TopLeft.Y+sprite.Height
	x1, y1 := sprite.TopLeft.X, sprite.TopLeft.Y
	xCenter, yCenter := x1+(x2-x1)/2, y1+(y2-y1)/2
	gc := draw2dimg.NewGraphicContext(image)
	gc.SetFillColor(color)

	if rotation == 0 {
		// TopLeft -> RightCenter
		gc.MoveTo(x1, y1)
		gc.LineTo(x2, yCenter)

		// RightCenter -> BottomRight
		gc.LineTo(x2, y2)

		// BottomRight -> BottomCenter
		gc.LineTo(xCenter, y2)

		// BottomCenter -> TopLeft
		gc.LineTo(x1, y1)
	} else if rotation == 1 {
		// TopRight -> BottomCenter
		gc.MoveTo(x2, y1)
		gc.LineTo(xCenter, y2)

		// BottomCenter -> BottomLeft
		gc.LineTo(x1, y2)

		// BottomLeft -> LeftCenter
		gc.LineTo(x1, yCenter)

		// LeftCenter -> TopRight
		gc.LineTo(x2, y1)
	} else if rotation == 2 {
		// TopLeft -> TopCenter
		gc.MoveTo(x1, y1)
		gc.LineTo(xCenter, y1)

		// TopCenter -> BottomRight
		gc.LineTo(x2, y2)

		// BottomRight -> LeftCenter
		gc.LineTo(x1, yCenter)

		// LeftCenter -> TopLeft
		gc.LineTo(x1, y1)
	} else {
		// TopCenter -> TopRight
		gc.MoveTo(xCenter, y1)
		gc.LineTo(x2, y1)

		// TopRight -> RightCenter
		gc.LineTo(x2, yCenter)

		// RightCenter -> BottomLeft
		gc.LineTo(x1, y2)

		// BottomLeft -> TopCenter
		gc.LineTo(xCenter, y1)
	}

	gc.SetLineWidth(0)

	// Fill
	gc.FillStroke()
}
