package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Type13Drawer struct {
	Name string
}

func (p *Type13Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int) {
	x2, y2 := sprite.TopLeft.X+sprite.Width, sprite.TopLeft.Y+sprite.Height
	x1, y1 := sprite.TopLeft.X, sprite.TopLeft.Y
	xCenter, yCenter := x1+sprite.Width/2, y1+sprite.Height/2
	gc := draw2dimg.NewGraphicContext(image)
	gc.SetFillColor(color)

	if rotation == 0 {
		// CenterCenter -> BottomRight
		gc.MoveTo(xCenter, yCenter)
		gc.LineTo(x2, y2)

		// BottomRight -> BottomLeft
		gc.MoveTo(x2, y2)
		gc.LineTo(x1, y2)

		// BottomLeft -> CenterCenter
		gc.MoveTo(x1, y2)
		gc.LineTo(xCenter, yCenter)
	} else if rotation == 1 {
		// TopLeft -> CenterCenter
		gc.MoveTo(x1, y1)
		gc.LineTo(xCenter, yCenter)

		// CenterCenter -> BottomLeft
		gc.MoveTo(xCenter, yCenter)
		gc.LineTo(x1, y2)

		// BottomLeft -> TopLeft
		gc.MoveTo(x1, y2)
		gc.LineTo(x1, y1)
	} else if rotation == 2 {
		// TopLeft -> TopRight
		gc.MoveTo(x1, y1)
		gc.LineTo(x2, y1)

		// TopRight -> CenterCenter
		gc.MoveTo(x2, y1)
		gc.LineTo(xCenter, yCenter)

		// CenterCenter -> TopLeft
		gc.MoveTo(xCenter, yCenter)
		gc.LineTo(x1, y1)
	} else {
		// TopRight -> CenterCenter
		gc.MoveTo(x2, y1)
		gc.LineTo(xCenter, yCenter)

		// CenterCenter -> BottomRight
		gc.MoveTo(xCenter, yCenter)
		gc.LineTo(x2, y2)

		// BottomRight -> TopRight
		gc.MoveTo(x2, y2)
		gc.LineTo(x2, y1)
	}

	gc.SetLineWidth(0)

	// Fill
	gc.FillStroke()
}
