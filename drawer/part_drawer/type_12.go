package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Type12Drawer struct {
	Name string
}

func (p *Type12Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int) {
	x2, y2 := sprite.TopLeft.X+sprite.Width, sprite.TopLeft.Y+sprite.Height
	x1, y1 := sprite.TopLeft.X, sprite.TopLeft.Y
	xCenter, yCenter := x1+sprite.Width/2, y1+sprite.Height/2
	gc := draw2dimg.NewGraphicContext(image)
	gc.SetFillColor(color)

	if rotation == 0 {
		// LeftCener -> RightCenter
		gc.MoveTo(x1, yCenter)
		gc.LineTo(x2, yCenter)

		// RightCenter -> BottomCenter
		gc.MoveTo(x2, yCenter)
		gc.LineTo(xCenter, y2)

		// BottomCenter -> LeftCenter
		gc.MoveTo(xCenter, y2)
		gc.LineTo(x1, yCenter)
	} else if rotation == 1 {
		// LeftCenter -> TopCenter
		gc.MoveTo(x1, yCenter)
		gc.LineTo(xCenter, y1)

		// TopCenter -> BottomCenter
		gc.MoveTo(xCenter, y1)
		gc.LineTo(xCenter, y2)

		// BottomCenter -> LeftCenter
		gc.MoveTo(xCenter, y2)
		gc.LineTo(x1, yCenter)
	} else if rotation == 2 {
		// TopCenter -> RigthCenter
		gc.MoveTo(xCenter, y1)
		gc.LineTo(x2, yCenter)

		// BottomCenter -> LeftCenter
		gc.MoveTo(x2, yCenter)
		gc.LineTo(x1, yCenter)

		// LeftCenter -> TopCenter
		gc.MoveTo(x1, yCenter)
		gc.LineTo(xCenter, y1)
	} else {
		// TopCenter -> RightCenter
		gc.MoveTo(xCenter, y1)
		gc.LineTo(x2, yCenter)

		// RightCenter -> BottomCenter
		gc.MoveTo(x2, yCenter)
		gc.LineTo(xCenter, y2)

		// BottomCenter -> TopCenter
		gc.MoveTo(xCenter, y2)
		gc.LineTo(xCenter, y1)
	}

	gc.SetLineWidth(0)

	// Fill
	gc.FillStroke()
}
