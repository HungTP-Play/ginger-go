package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Type10Drawer struct {
	Name string
}

func (p *Type10Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int) {
	x2, y2 := sprite.TopLeft.X+sprite.Width, sprite.TopLeft.Y+sprite.Height
	x1, y1 := sprite.TopLeft.X, sprite.TopLeft.Y
	xCenter, yCenter := x1+sprite.Width/2, y1+sprite.Height/2
	gc := draw2dimg.NewGraphicContext(image)
	gc.SetFillColor(color)

	if rotation == 0 {
		// TopCenter -> TopRight
		gc.MoveTo(xCenter, y1)
		gc.LineTo(x2, y1)

		// TopRight -> BottomLeft
		gc.LineTo(x1, y2)

		// BottomLeft -> LeftCenter
		gc.LineTo(x1, yCenter)

		// LeftCenter -> CenterCenter
		gc.LineTo(xCenter, yCenter)

		// CenterCenter -> TopCenter
		gc.LineTo(xCenter, y1)
	} else if rotation == 1 {
		// TopLeft -> TopCenter
		gc.MoveTo(x1, y1)
		gc.LineTo(xCenter, y1)

		// TopCenter -> CenterCenter
		gc.LineTo(xCenter, yCenter)

		// CenterCenter -> RightCenter
		gc.LineTo(x2, yCenter)

		// RightCenter -> BottomRight
		gc.LineTo(x2, y2)

		// BottomRight -> TopLeft
		gc.LineTo(x1, y1)
	} else if rotation == 2 {
		// TopRight -> RightCenter
		gc.MoveTo(x2, y1)
		gc.LineTo(x2, yCenter)

		// RightCenter -> CenterCenter
		gc.LineTo(xCenter, yCenter)

		// CenterCenter -> BottomCenter
		gc.LineTo(xCenter, y2)

		// BottomCenter -> BottomLeft
		gc.LineTo(x1, y2)

		// BottomLeft -> TopRight
		gc.LineTo(x2, y1)
	} else {
		// BottomRight -> BottomCenter
		gc.MoveTo(x2, y2)
		gc.LineTo(xCenter, y2)

		// BottomCenter -> CenterCenter
		gc.LineTo(xCenter, yCenter)

		// CenterCenter -> LeftCenter
		gc.LineTo(x1, yCenter)

		// LeftCenter -> TopLeft
		gc.LineTo(x1, y1)

		// TopLeft -> BottomRight
		gc.LineTo(x2, y2)
	}

	gc.SetLineWidth(0)

	// Fill
	gc.FillStroke()
}
