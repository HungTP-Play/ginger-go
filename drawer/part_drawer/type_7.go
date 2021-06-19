package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Type7Drawer struct {
	Name string
}

func (p *Type7Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int) {
	x2, y2 := sprite.TopLeft.X+sprite.Width, sprite.TopLeft.Y+sprite.Height
	x1, y1 := sprite.TopLeft.X, sprite.TopLeft.Y
	xCenter, yCenter := x1+sprite.Width/2, y1+sprite.Height/2
	x1p4, x3p4 := x1+sprite.Width*0.25, x1+sprite.Width*0.75
	y1p4, y3p4 := y1+sprite.Height*0.25, y1+sprite.Height*0.75
	gc := draw2dimg.NewGraphicContext(image)

	if rotation == 0 {
		// TopCenter -> BottomRight
		gc.MoveTo(xCenter, y1)
		gc.LineTo(x2, y2)

		// BottomRight -> BottomCenter
		gc.MoveTo(x2, y2)
		gc.LineTo(xCenter, y2)

		// BottomCenter -> Center3p4
		gc.MoveTo(xCenter, y2)
		gc.LineTo(x3p4, yCenter)

		// Center3p4 -> Center1p4
		gc.MoveTo(x3p4, yCenter)
		gc.LineTo(x1p4, yCenter)

		// Center1p4 -> BottomCenter
		gc.MoveTo(x1p4, yCenter)
		gc.LineTo(xCenter, y2)

		// BottomCenter -> BottomLeft
		gc.MoveTo(xCenter, y2)
		gc.LineTo(x1, y2)

		// BottomLeft -> TopCenter
		gc.MoveTo(x1, y2)
		gc.LineTo(xCenter, y1)
	} else if rotation == 1 {
		// TopLeft -> RightCenter
		gc.MoveTo(x1, y1)
		gc.LineTo(x2, yCenter)

		// RightCenter -> BottomLeft
		gc.MoveTo(x2, yCenter)
		gc.LineTo(x1, y2)

		// BottomLeft -> LeftCenter
		gc.MoveTo(x1, y2)
		gc.LineTo(x1, yCenter)

		// LeftCenter -> xCentery34
		gc.MoveTo(x1, yCenter)
		gc.LineTo(xCenter, y3p4)

		// xCentery34 -> xCentery14
		gc.MoveTo(xCenter, y3p4)
		gc.LineTo(xCenter, y1p4)

		// xCentery14 -> LeftCenter
		gc.MoveTo(xCenter, y1p4)
		gc.LineTo(x1, yCenter)

		// LeftCenter -> TopLeft
		gc.MoveTo(x1, yCenter)
		gc.LineTo(x1, y1)
	} else if rotation == 2 {
		// TopLeft -> TopRight
		gc.MoveTo(x1, y1)
		gc.LineTo(x2, y1)

		// TopRight -> BottomCenter
		gc.MoveTo(x2, y1)
		gc.LineTo(xCenter, y2)

		// BottomCenter -> X14yCenter
		gc.MoveTo(xCenter, y2)
		gc.LineTo(x1p4, yCenter)

		// X14yCenter -> X34YCenter
		gc.MoveTo(x1p4, yCenter)
		gc.LineTo(x3p4, yCenter)

		// X34YCenter -> TopCenter
		gc.MoveTo(x3p4, yCenter)
		gc.LineTo(xCenter, y1)

		// TopCenter -> X14YCenter
		gc.MoveTo(xCenter, y1)
		gc.LineTo(x1p4, yCenter)

		// X14YCenter -> TopLeft
		gc.MoveTo(x1p4, yCenter)
		gc.LineTo(x1, y1)
	} else {
		// TopRight -> LeftCenter
		gc.MoveTo(x2, y1)
		gc.LineTo(x1, yCenter)

		// LeftCenter -> BottomRight
		gc.MoveTo(x1, yCenter)
		gc.LineTo(x2, y2)

		// BottomRight -> RightCenter
		gc.MoveTo(x2, y2)
		gc.LineTo(x2, yCenter)

		// RightCenter -> CenterY24
		gc.MoveTo(x2, yCenter)
		gc.LineTo(xCenter, y3p4)

		// CenterY34 -> CenterY14
		gc.MoveTo(xCenter, y3p4)
		gc.LineTo(xCenter, y1p4)

		// CenterY14 -> RightCenter
		gc.MoveTo(xCenter, y1p4)
		gc.LineTo(x2, yCenter)

		// RightCenter -> TopRight
		gc.MoveTo(x2, yCenter)
		gc.LineTo(x2, y1)
	}

	gc.SetLineWidth(0)

	// Fill
	gc.FillStroke()
}
