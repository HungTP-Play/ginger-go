package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Type8Drawer struct {
	Name string
}

func (p *Type8Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int) {
	x2, y2 := sprite.TopLeft.X+sprite.Width, sprite.TopLeft.Y+sprite.Height
	x1, y1 := sprite.TopLeft.X, sprite.TopLeft.Y
	xCenter, yCenter := x1+sprite.Width/2, y1+sprite.Height/2
	gc := draw2dimg.NewGraphicContext(image)

	if rotation == 0 {
		// TopLeft -> RightCenter
		gc.MoveTo(x1, y1)
		gc.LineTo(x2, yCenter)

		// RightCenter -> BottomCenter
		gc.MoveTo(x2, yCenter)
		gc.LineTo(xCenter, y2)

		// BottomCenter -> TopLeft
		gc.MoveTo(xCenter, y2)
		gc.LineTo(x1, y1)
	} else if rotation == 1 {
		// TopRight -> LeftCenter
		gc.MoveTo(x2, y1)
		gc.LineTo(x1, yCenter)

		// LeftCenter -> BottomCenter
		gc.MoveTo(x1, yCenter)
		gc.LineTo(xCenter, y2)

		// BottomCenter -> TopRight
		gc.MoveTo(xCenter, y2)
		gc.LineTo(x2, y1)
	} else if rotation == 2 {
		// TopCenter -> LeftCenter
		gc.MoveTo(xCenter, y1)
		gc.LineTo(x1, yCenter)

		// LeftCenter -> BottomRight
		gc.MoveTo(x1, yCenter)
		gc.LineTo(x2, y2)

		// BottomRight -> TopCenter
		gc.MoveTo(x2, y2)
		gc.LineTo(xCenter, y1)
	} else {
		// TopCenter -> RightCenter
		gc.MoveTo(xCenter, y1)
		gc.LineTo(x2, yCenter)

		// RightCenter -> BottomLeft
		gc.MoveTo(x2, yCenter)
		gc.LineTo(x1, y2)

		// BottomLeft -> TopCenter
		gc.MoveTo(x1, y2)
		gc.LineTo(xCenter, y1)
	}

	gc.SetLineWidth(0)

	// Fill
	gc.FillStroke()
}
