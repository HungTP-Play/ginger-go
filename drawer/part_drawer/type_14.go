package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Type14Drawer struct {
	Name string
}

func (p *Type14Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int) {
	x2, y2 := sprite.TopLeft.X+sprite.Width, sprite.TopLeft.Y+sprite.Height
	x1, y1 := sprite.TopLeft.X, sprite.TopLeft.Y
	xCenter, yCenter := x1+sprite.Width/2, y1+sprite.Height/2
	gc := draw2dimg.NewGraphicContext(image)

	if rotation == 0 {
		// TopCenter -> LeftCenter
		gc.MoveTo(xCenter, y1)
		gc.LineTo(x1, yCenter)

		// LeftCenter -> CenterCenter
		gc.MoveTo(x1, yCenter)
		gc.LineTo(xCenter, yCenter)

		// CenterCenter -> TopCenter
		gc.MoveTo(xCenter, yCenter)
		gc.LineTo(xCenter, y1)
	} else if rotation == 1 {
		// TopRight -> CenterCenter
		gc.MoveTo(x2, y1)
		gc.LineTo(xCenter, yCenter)

		// CenterCenter -> RightCenter
		gc.MoveTo(xCenter, yCenter)
		gc.LineTo(x2, yCenter)

		// RightCenter -> TopRight
		gc.MoveTo(x2, yCenter)
		gc.LineTo(x2, y1)
	} else if rotation == 2 {
		// CenterCenter -> RightCenter
		gc.MoveTo(xCenter, yCenter)
		gc.LineTo(x2, yCenter)

		// RightCenter -> BottomCenter
		gc.MoveTo(x2, yCenter)
		gc.LineTo(xCenter, y2)

		// BottomCenter -> CenterCenter
		gc.MoveTo(xCenter, y2)
		gc.LineTo(xCenter, yCenter)
	} else {
		// LeftCenter -> CenterCenter
		gc.MoveTo(x1, yCenter)
		gc.LineTo(xCenter, yCenter)

		// CenterCenter -> BottomCenter
		gc.MoveTo(xCenter, yCenter)
		gc.LineTo(xCenter, y2)

		// BottomCenter -> LeftCenter
		gc.MoveTo(xCenter, y2)
		gc.LineTo(x1, yCenter)
	}

	gc.SetLineWidth(0)

	// Fill
	gc.FillStroke()
}